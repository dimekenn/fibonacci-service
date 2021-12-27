package service

import (
	"context"
	"encoding/json"
	"fibonacciService/internal/app/models"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"math/big"
	"net/http"
	"time"
)

type ServiceImpl struct {
	redisCli *redis.Client
}

func NewService(redisCli *redis.Client) Service {
	return &ServiceImpl{redisCli: redisCli}
}

func (s ServiceImpl) GetFibonacciList(ctx context.Context, from, to uint64) (*models.FibResponse, error) {
	//if x>y or x=y returns error
	if from > to || from == to {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "x should be more than y")
	}
	//copy of fibresponse object
	res := &models.FibResponse{}

	//trying to find cache in redis
	redisRes, redisErr := s.redisCli.Get(ctx, fmt.Sprintf("%d,%d", from, to)).Result()
	if redisErr == nil {
		log.Infof("results found in redis cache")
		//cache found in redis and deserializing to ibject
		unmErr := json.Unmarshal([]byte(redisRes), res)
		if unmErr != nil {
			log.Errorf("failed to unmarshall result from redis: %v", unmErr)
			return calculateAndCache(from, to, res, ctx, s.redisCli), nil
		}
		return res, nil
	}
	//key by x and y is empty and calculating fibonacci numbers
	log.Infof("calculating fibonacci")
	return calculateAndCache(from, to, res, ctx, s.redisCli), nil
}

//function for calculate fibonacci by given arguments
func calculateAndCache(from, to uint64, res *models.FibResponse, ctx context.Context, redisCli *redis.Client) *models.FibResponse {
	//created func for calculate fibonacci numbers
	f := fib(from)
	var i uint64
	//created loop by argument y
	for i = 0; i < to; i++ {
		res.FibonacciSlice = append(res.FibonacciSlice, f().String())
	}

	//serializing slice to json object to save in cache
	marRes, marErr := json.Marshal(res)
	if marErr != nil {
		log.Errorf("failed to marshall res: %v", marErr)
	}

	//saving data in redis cache
	redisSetErr := redisCli.Set(ctx, fmt.Sprintf("%d,%d", from, to), marRes, 24*time.Hour).Err()
	if redisSetErr != nil {
		log.Errorf("failed to set to redis cli: %v", redisSetErr)
	}
	return res
}

//calculator for fibonacci numbers
func fib(from uint64) func() *big.Int {
	x := big.NewInt(int64(from))
	y := big.NewInt(int64(from))
	y = y.Add(y, big.NewInt(1))

	return func() *big.Int {
		x, y = y, x.Add(x, y)
		return x
	}
}

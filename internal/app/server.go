package app

import (
	"fibonacciService/configs"
	"fibonacciService/internal/app/handler"
	"fibonacciService/internal/app/service"
	"fibonacciService/proto"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net"
)

//server for REST apis
func StartHTTPServer(errCh chan<- error, cfg *configs.Configs) {
	app := echo.New()

	fibService := service.NewService(startRedisClient(cfg))
	fibHandler := handler.NewHandler(fibService)

	app.GET("/api/v1/fibonacci", fibHandler.GetFibonacci)

	errCh <- app.Start(cfg.Port)
}

//server for GRPC apis
func StartGRPCServer(errCh chan<- error, cfg *configs.Configs){
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil{
		errCh <- err
	}
	s:= grpc.NewServer()
	fibService := service.NewService(startRedisClient(cfg))
	fibHandler := handler.NewHandler(fibService)
	proto.RegisterFibonacciServiceServer(s, fibHandler)

	log.Infof("grpc server started on :%s", cfg.GRPCPort)

	errCh <- s.Serve(lis)
}

//new client for Redis
func startRedisClient(cfg *configs.Configs) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
		Password: "",
		DB: 0,
	})
}


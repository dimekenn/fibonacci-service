package service

import (
	"fibonacciService/internal/app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ServiceImpl struct {

}

func NewService() Service {
	return &ServiceImpl{}
}

func (s ServiceImpl) GetFibonacciList(from uint64, to uint64) (*models.FibResponse, error) {
	if to > 95{
		return nil, echo.NewHTTPError(http.StatusBadRequest, "too large value for y")
	} else if from > to{
		return nil, echo.NewHTTPError(http.StatusBadRequest, "x should be more than y")
	}

	f := fib(from)
	res := &models.FibResponse{}
	var i uint64
	for i = 0; i < to; i++{
		res.FibonacciSlice = append(res.FibonacciSlice, f())
	}

	return res, nil
}

func fib(from uint64) func() uint64{
	x := from
	y := from+1

	return func() uint64 {
		x, y = y, x + y
		return x
	}
}
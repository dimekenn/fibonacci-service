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

func (s ServiceImpl) GetFibonacciList(from int, to int) (*models.FibResponse, error) {
	if to > 95{
		return nil, echo.NewHTTPError(http.StatusBadRequest, "too large value for y")
	} else if from > to{
		return nil, echo.NewHTTPError(http.StatusBadRequest, "x should be more than y")
	}

	f := fib(from)
	res := &models.FibResponse{}
	for i := 0; i < to; i++{
		res.FibonacciSlice = append(res.FibonacciSlice, f())
	}

	return res, nil
}

func fib(from int) func() int{
	x := from
	y := from+1

	return func() int {
		x, y = y, x + y
		return x
	}
}
package service

import (
	"context"
	"fibonacciService/internal/app/models"
)

type Service interface {
	GetFibonacciList(ctx context.Context, from, to uint64) (*models.FibResponse, error)
}
package service

import "fibonacciService/internal/app/models"

type Service interface {
	GetFibonacciList(from uint64, to uint64) (*models.FibResponse, error)
}
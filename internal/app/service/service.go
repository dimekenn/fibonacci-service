package service

import "fibonacciService/internal/app/models"

type Service interface {
	GetFibonacciList(from int, to int) (*models.FibResponse, error)
}
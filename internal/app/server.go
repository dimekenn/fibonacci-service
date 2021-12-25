package app

import (
	"fibonacciService/configs"
	"fibonacciService/internal/app/handler"
	"fibonacciService/internal/app/service"
	"github.com/labstack/echo/v4"
)

func StartHTTPServer(errCh chan<- error, cfg *configs.Configs) {
	app := echo.New()

	fibService := service.NewService()
	fibHandler := handler.NewHandler(fibService)

	app.GET("/api/v1/fibonacci", fibHandler.GetFibonacci)

	errCh <- app.Start(cfg.Port)
}


package handler

import (
	"fibonacciService/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetFibonacci(c echo.Context) error {
	x := c.QueryParam("x")
	y := c.QueryParam("y")
	from, err := strconv.Atoi(x)
	to, err := strconv.Atoi(y)
	if err != nil{
		log.Errorf("failed to parse params: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot parse params")
	}
	res, serErr := h.service.GetFibonacciList(from, to)
	if serErr != nil{
		log.Errorf("failed to calculate fibonacci: %v", serErr)
		return serErr
	}
	log.Infof("Success response: %v", res)
	return c.JSON(http.StatusOK, res)
}

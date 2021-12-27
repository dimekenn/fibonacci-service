package handler

import (
	"context"
	"fibonacciService/internal/app/service"
	"fibonacciService/proto"
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
	log.Infof("new request x = %d, y = %d", from, to)
	res, serErr := h.service.GetFibonacciList(c.Request().Context(), uint64(from), uint64(to))
	if serErr != nil{
		log.Errorf("failed to calculate fibonacci: %v", serErr)
		return serErr
	}
	log.Infof("Success response: %v", res)
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) GetFibonacciSlice(ctx context.Context, req *proto.GetFibonacciSliceReq) (*proto.GetFibonacciSliceRes, error) {
	log.Infof("new request on GRPC server x = %d, y = %d", req.X, req.Y)
	fibSlice, err := h.service.GetFibonacciList(ctx, req.X, req.Y)
	if err != nil{
		log.Errorf("failed to calculate fibonacci on GRPC server: %v", err)
		return nil, err
	}
	res:=&proto.GetFibonacciSliceRes{Res: fibSlice.FibonacciSlice}
	log.Infof("Success response on GRPC server: %v", fibSlice)
	return res, nil
}


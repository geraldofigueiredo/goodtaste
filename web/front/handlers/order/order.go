package order

import (
	"github.com/geraldofigueiredo/goodtaste/web/front/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type orderHandler struct {
	orderSvc service.OrderService
	logger   *zap.SugaredLogger
}

func NewOrderHandler(e *echo.Echo, orderSvc service.OrderService, logger *zap.SugaredLogger) {
	h := orderHandler{
		orderSvc: orderSvc,
		logger:   logger,
	}

	e.POST("/order", h.NewOrder)
}

func (h *orderHandler) NewOrder(c echo.Context) error {
	return c.JSONBlob(http.StatusCreated, []byte("Ok"))
}

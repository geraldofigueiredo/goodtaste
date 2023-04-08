package handlers

import (
	"github.com/geraldofigueiredo/goodtaste/web/front/handlers/order"
	"github.com/geraldofigueiredo/goodtaste/web/front/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func RegisterHandlers(e *echo.Echo, orderSvc service.OrderService, logger *zap.SugaredLogger) {
	e.GET("/", handlerOK)

	order.NewOrderHandler(e, orderSvc, logger)
}

func handlerOK(c echo.Context) error {
	return c.String(http.StatusOK, "Ok")
}

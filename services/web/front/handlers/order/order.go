package order

import (
	"github.com/geraldofigueiredo/goodtaste/services/web/front/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
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
	e.GET("/order/:id", h.GetOrder)
}

func (h *orderHandler) NewOrder(c echo.Context) error {
	order := newOrderRequest{}

	err := c.Bind(&order)
	if err != nil {
		return err
	}

	err = h.orderSvc.NewOrder()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSONBlob(http.StatusCreated, []byte("Ok"))
}

func (h *orderHandler) GetOrder(c echo.Context) error {
	id := c.Param("id")

	order, err := h.orderSvc.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, &orderInfo{
		ID:        order.ID,
		OwnerName: order.OwnerName,
		Infos:     order.Infos,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	})
}

type newOrderRequest struct {
	Infos     []string `json:"infos"`
	OwnerName string   `json:"owner_name"`
}

type orderInfo struct {
	ID        uuid.UUID  `json:"id"`
	OwnerName string     `json:"owner_name"`
	Infos     []string   `json:"infos"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

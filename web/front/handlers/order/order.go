package order

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/geraldofigueiredo/goodtaste/web/front/service"
	"net/http"
)

type orderHandler struct {
	orderSvc service.OrderService
}

func NewOrderHandler(ws *restful.WebService, orderSvc service.OrderService) {
	h := orderHandler{
		orderSvc: orderSvc,
	}

	ws.Route(ws.POST("/order").
		To(h.NewOrder),
	)

	ws.Route(ws.GET("/order").
		To(h.NewOrder2),
	)
}

func (h *orderHandler) NewOrder(request *restful.Request, response *restful.Response) {
	fmt.Println("newOrder - POST")
	response.WriteHeader(http.StatusCreated)
}

func (h *orderHandler) NewOrder2(request *restful.Request, response *restful.Response) {
	fmt.Println("newOrder - GET")
	response.WriteHeader(http.StatusOK)
}

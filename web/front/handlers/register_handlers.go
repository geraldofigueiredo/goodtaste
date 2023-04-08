package handlers

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/geraldofigueiredo/goodtaste/web/front/handlers/order"
	"github.com/geraldofigueiredo/goodtaste/web/front/service"
)

func RegisterHandlers(container *restful.Container, orderSvc service.OrderService) {
	ws := new(restful.WebService)

	ws.Route(ws.GET("/").To(handlerOK))
	//ws.Consumes("application/json").Produces("application/json")

	order.NewOrderHandler(ws, orderSvc)

	container.Add(ws)
}

func handlerOK(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK"))
}

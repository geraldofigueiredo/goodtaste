package main

import (
	"fmt"
	restful "github.com/emicklei/go-restful/v3"
	"github.com/geraldofigueiredo/goodtaste/web/front/handlers"
	"github.com/geraldofigueiredo/goodtaste/web/front/service"
	"net/http"
)

func main() {
	orderSvc := service.NewOrderService()

	handlers.RegisterHandlers(restful.DefaultContainer, orderSvc)

	fmt.Println("listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

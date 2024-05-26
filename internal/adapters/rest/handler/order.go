package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
)

type OrderHandler struct {
	service *order.Service
}

func RegisterOrderHandlers(router *gin.Engine, service *order.Service) {
	orderController := controller.NewOrderController(service)

	router.GET("/order", orderController.GetAll)
	router.GET("/order/:id", orderController.Get)
}

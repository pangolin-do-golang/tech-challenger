package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/order"
)

type OrderHandler struct {
	service *cart.Service
}

func RegisterOrderHandlers(router *gin.Engine, service order.IOrderService) {
	orderController := controller.NewOrderController(service)

	// todo change to orders
	router.POST("/order", orderController.Create)
	router.GET("/order", orderController.GetAll)
	router.GET("/order/:id", orderController.Get)
	router.POST("/orders/webhook", orderController.Webhook)
}

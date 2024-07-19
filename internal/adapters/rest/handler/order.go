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

	router.POST("/orders", orderController.Create)
	router.GET("/orders", orderController.GetAll)
	router.GET("/orders/:id", orderController.Get)
	router.PATCH("/orders/:id", orderController.Update)
}

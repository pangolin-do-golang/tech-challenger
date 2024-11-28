package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/core/customer"
)

func RegisterCustomerHandlers(router *gin.Engine, service customer.IService) {
	customerController := controller.NewCustomerController(service)

	router.POST("/customer", customerController.Create)
	router.PUT("/customer/:id", customerController.Update)
	router.DELETE("/customer/:id", customerController.Delete)
	router.GET("/customer", customerController.GetAll)
	router.GET("/customer/:cpf", customerController.GetByCpf)
}

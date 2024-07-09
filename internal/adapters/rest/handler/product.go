package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
)

func RegisterProductHandlers(router *gin.Engine, service *cart.Service) {
	productController := controller.NewProductController(service)

	router.GET("/product", productController.Search)

	router.DELETE("/product/:id")
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
)

func RegisterCartHandlers(router *gin.Engine, service cart.IService) {
	cartController := controller.NewCartController(service)

	router.POST("/cart/add-product", cartController.AddProduct)
	router.POST("/cart/remove-product", cartController.RemoveProduct)
	router.POST("/cart/edit-product", cartController.EditProduct)
}

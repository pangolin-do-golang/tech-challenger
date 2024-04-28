package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"net/http"
)

func RegisterCartHandlers(router *gin.Engine, service cart.IService) {
	router.POST("/cart/add-product", func(c *gin.Context) {
		type Payload struct {
			ClientID  string `json:"client_id" binding:"required"`
			ProductID string `json:"product_id" binding:"required"`
			// TODO validação de quantidade >= 0 / estoque
			Quantity int    `json:"quantity" binding:"required"`
			Comments string `json:"comments"`
		}

		payload := &Payload{}
		err := c.BindJSON(payload)
		if err != nil {
			// TODO melhorar retorno de validação
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.AddProduct(c.Request.Context(), &cart.Product{
			ClientID:  payload.ClientID,
			ProductID: payload.ProductID,
			Quantity:  payload.Quantity,
			Comments:  payload.Comments,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.Status(http.StatusOK)
	})

	router.POST("/cart/remove-product", func(c *gin.Context) {
		type Payload struct {
			ClientID  string `json:"client_id" binding:"required"`
			ProductID string `json:"product_id" binding:"required"`
		}
		payload := &Payload{}
		err := c.BindJSON(payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.RemoveProduct(c.Request.Context(), payload.ClientID, payload.ProductID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.Status(http.StatusOK)
	})

	router.POST("/cart/edit-product", func(c *gin.Context) {
		type Payload struct {
			ClientID  string `json:"client_id" binding:"required"`
			ProductID string `json:"product_id" binding:"required"`
			Quantity  int    `json:"quantity" binding:"required"`
			Comments  string `json:"comments"`
		}
		payload := &Payload{}
		err := c.BindJSON(payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.EditProduct(c.Request.Context(), &cart.Product{
			ClientID:  payload.ClientID,
			ProductID: payload.ProductID,
			Quantity:  payload.Quantity,
			Comments:  payload.Comments,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.Status(http.StatusOK)
	})
}

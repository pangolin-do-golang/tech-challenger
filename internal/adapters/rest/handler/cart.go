package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"net/http"
)

type CartAddProductPayload struct {
	ClientID  string `json:"client_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
	Comments  string `json:"comments"`
}

func RegisterCartHandlers(router *gin.Engine, service cart.IService) {
	router.POST("/cart/add-product", func(c *gin.Context) {
		payload := &CartAddProductPayload{}
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

		c.Status(http.StatusCreated)
	})
}

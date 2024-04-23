package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"net/http"
)

type CartAddProductPayload struct {
	ClientID  string `json:"client_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Comments  string `json:"comments"`
}

func RegisterCartHandlers(router *gin.Engine, service cart.IService) {
	router.POST("/cart/add-product", func(c *gin.Context) {
		var payload *CartAddProductPayload
		err := c.BindJSON(payload)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		err = service.AddProduct(c.Request.Context(), &cart.Product{
			ClientID:  payload.ClientID,
			ProductID: payload.ProductID,
			Quantity:  payload.Quantity,
			Comments:  payload.Comments,
		})

		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		c.Status(http.StatusOK)
	})
}

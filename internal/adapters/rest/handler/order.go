package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"net/http"
)

func RegisterOrderHandlers(router *gin.Engine, service *order.Service) {

	router.GET("/order/:id", func(c *gin.Context) {
		id := c.Param("id")

		order, err := service.Get(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went bad :(",
			})

			return
		}

		c.JSON(http.StatusOK, order)
	})
}

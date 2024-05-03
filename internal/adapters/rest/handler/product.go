package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/product"
	"net/http"
)

func RegisterProductHandlers(router *gin.Engine, service *product.Service) {

	router.GET("/product", func(c *gin.Context) {
		search := c.Query("search")
		category := c.Query("category")

		products, err := service.Search(search, category)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went bad :(",
			})

			return
		}

		c.JSON(http.StatusOK, products)
	})

	router.DELETE("/product/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := service.Delete(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went bad :(",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})
}

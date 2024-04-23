package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/product"
	"net/http"
)

type RestServer struct {
	orderService   *order.Service
	productService *product.Service
}

func NewRestServer(orderService *order.Service, productService *product.Service) *RestServer {
	return &RestServer{
		orderService:   orderService,
		productService: productService,
	}
}

func (rs RestServer) Serve() {
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	handler.RegisterOrderHandlers(r, rs.orderService)
	handler.RegisterProductHandlers(r, rs.productService)

	r.Run()
}

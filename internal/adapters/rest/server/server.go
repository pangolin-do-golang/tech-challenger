package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"net/http"
)

type RestServer struct {
	orderService *order.Service
}

func NewRestServer(orderService *order.Service) *RestServer {
	return &RestServer{
		orderService: orderService,
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

	r.Run()
}

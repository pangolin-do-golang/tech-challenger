package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/customer"
)

type RestServer struct {
	customerService customer.IService
}

type RestServerOptions struct {
	CustomerService customer.IService
}

func NewRestServer(options *RestServerOptions) *RestServer {
	return &RestServer{
		customerService: options.CustomerService,
	}
}

func (rs RestServer) Serve() {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	handler.RegisterCustomerHandlers(r, rs.customerService)
	handler.RegisterSwaggerHandlers(r)
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

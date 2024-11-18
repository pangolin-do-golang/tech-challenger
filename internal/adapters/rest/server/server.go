package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/product"
)

type RestServer struct {
	productService *product.Service
	cartService    cart.IService
}

type RestServerOptions struct {
	ProductService *product.Service
	CartService    cart.IService
}

func NewRestServer(options *RestServerOptions) *RestServer {
	return &RestServer{
		productService: options.ProductService,
		cartService:    options.CartService,
	}
}

func (rs RestServer) Serve() {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	handler.RegisterProductHandlers(r, rs.productService)
	handler.RegisterCartHandlers(r, rs.cartService)
	handler.RegisterSwaggerHandlers(r)
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}

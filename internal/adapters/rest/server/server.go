package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/product"
	"net/http"
)

type RestServer struct {
	orderService   *order.Service
	productService *product.Service
	cartService    cart.IService
}

type RestServerOptions struct {
	OrderService   *order.Service
	ProductService *product.Service
	CartService    cart.IService
}

func NewRestServer(options *RestServerOptions) *RestServer {
	return &RestServer{
		orderService:   options.OrderService,
		productService: options.ProductService,
		cartService:    options.CartService,
	}
}

func (rs RestServer) Serve() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	handler.RegisterOrderHandlers(r, rs.orderService)
	handler.RegisterProductHandlers(r, rs.productService)
	handler.RegisterCartHandlers(r, rs.cartService)

	r.Run()
}

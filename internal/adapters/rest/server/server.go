package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/customer"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
)

type RestServer struct {
	orderService    order.IOrderService
	productService  *cart.Service
	cartService     cart.IService
	customerService customer.IService
}

type RestServerOptions struct {
	OrderService    order.IOrderService
	ProductService  *cart.Service
	CartService     cart.IService
	CustomerService customer.IService
}

func NewRestServer(options *RestServerOptions) *RestServer {
	return &RestServer{
		orderService:    options.OrderService,
		productService:  options.ProductService,
		cartService:     options.CartService,
		customerService: options.CustomerService,
	}
}

func (rs RestServer) Serve() {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	handler.RegisterOrderHandlers(r, rs.orderService)
	handler.RegisterProductHandlers(r, rs.productService)
	handler.RegisterCartHandlers(r, rs.cartService)
	handler.RegisterCustomerHandlers(r, rs.customerService)
	handler.RegisterSwaggerHandlers(r)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

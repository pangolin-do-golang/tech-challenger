package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/pangolin-do-golang/tech-challenge/docs"
	dbAdapter "github.com/pangolin-do-golang/tech-challenge/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// @title Tech Challenge Food API
// @version 0.1.0
// @description Fast Food API for FIAP Tech course

// @host localhost:8080
// @BasePath /
func main() {
	db, err := initDb()
	if err != nil {
		panic(err)
	}
	orderRepository := dbAdapter.NewPostgresOrderRepository(db)
	orderService := order.NewOrderService(orderRepository)

	productRepository := dbAdapter.NewPostgresProductRepository(db)
	productService := product.NewProductService(productRepository)

	cartRepository := dbAdapter.NewPostgresCartRepository(db)
	cartProductsRepository := dbAdapter.NewPostgresCartProductsRepository(db)
	cartService := cart.NewService(cartRepository, cartProductsRepository)

	restServer := server.NewRestServer(&server.RestServerOptions{
		OrderService:   orderService,
		ProductService: productService,
		CartService:    cartService,
	})

	restServer.Serve()
}

func initDb() (*gorm.DB, error) {
	_ = godotenv.Load()
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return db, nil
}

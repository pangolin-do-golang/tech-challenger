package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/pangolin-do-golang/tech-challenge/docs"
	dbAdapter "github.com/pangolin-do-golang/tech-challenge/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/cart"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/customer"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/order"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	customerService := customer.NewService()

	productRepository := dbAdapter.NewPostgresProductRepository(db)
	productService := product.NewProductService(productRepository)

	cartService := cart.NewService()

	orderService := order.NewOrderService(cartService, productService)

	restServer := server.NewRestServer(&server.RestServerOptions{
		OrderService:    orderService,
		ProductService:  productService,
		CartService:     cartService,
		CustomerService: customerService,
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

	err = db.AutoMigrate(
		&dbAdapter.CustomerPostgres{},
		&dbAdapter.ProductPostgres{},
		&dbAdapter.OrderPostgres{},
		&dbAdapter.CartPostgres{},
		&dbAdapter.CartProductsPostgres{},
		&dbAdapter.OrderPostgres{},
		&dbAdapter.OrderProductPostgres{},
	)
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}

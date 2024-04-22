package main

import (
	"fmt"
	db2 "github.com/pangolin-do-golang/tech-challenge/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func main() {
	db, err := initDb()
	if err != nil {
		panic(err)
	}
	orderRepository := db2.NewPostgresOrderRepository(db)
	orderService := order.NewOrderService(orderRepository)
	restServer := server.NewRestServer(orderService)
	restServer.Serve()
}

func initDb() (*gorm.DB, error) {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		loc,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return db, nil
}

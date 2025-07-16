package main

import (
	"fashiondemon/internal/config"
	"fashiondemon/internal/order"
	"fashiondemon/internal/product"
	"fashiondemon/internal/user"
	"net/http"

	"gorm.io/gorm"
)

func main() {
	config.InitDB()
	runMigrations(config.DB)

	mux := http.NewServeMux()
	user.RegisterRoutes(mux)
	product.RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&product.Category{},
		&product.Product{},
		&user.User{},
		&order.Order{},
		&order.OrderItem{},
	)
}

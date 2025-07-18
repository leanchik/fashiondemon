package main

import (
	"fashiondemon/internal/config"
	"fashiondemon/internal/order"
	"fashiondemon/internal/product"
	"fashiondemon/internal/user"
	"net/http"
)

func main() {
	config.InitDB()

	// основные миграции (без order)
	runMigrations()

	// миграция заказов (внутри order пакета)
	order.Migrate()

	// роуты
	mux := http.NewServeMux()
	user.RegisterRoutes(mux)
	product.RegisterRoutes(mux)
	order.RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}

func runMigrations() {
	config.DB.AutoMigrate(
		&product.Category{},
		&product.Product{},
		&user.User{},
	)
}

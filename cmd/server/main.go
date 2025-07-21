// @title FashionDemon API
// @version 1.0
// @description Документация API для магазина одежды FashionDemon

// @host localhost:8080
// @BasePath /

package main

import (
	_ "fashiondemon/docs"
	"fashiondemon/internal/config"
	"fashiondemon/internal/order"
	"fashiondemon/internal/product"
	"fashiondemon/internal/user"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func main() {
	config.InitDB()

	runMigrations()

	order.Migrate()

	mux := http.NewServeMux()
	user.RegisterRoutes(mux)
	product.RegisterRoutes(mux)
	order.RegisterRoutes(mux)
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs"))))
	mux.HandleFunc("/swagger/index.html", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", mux)
}

func runMigrations() {
	config.DB.AutoMigrate(
		&product.Category{},
		&product.Product{},
		&user.User{},
	)
}

package product

import (
	"fashiondemon/pkg/auth"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	adminOnly := http.HandlerFunc(CreateProductHandler)
	mux.Handle("/admin/product", auth.JWTMiddleware(adminOnly))
	mux.Handle("/admin/category/", auth.JWTMiddleware(http.HandlerFunc(CreateCategoryHandler)))

	mux.HandleFunc("/products", GetAllProducts)
}

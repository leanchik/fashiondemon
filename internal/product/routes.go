package product

import (
	"fashiondemon/pkg/auth"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	adminOnly := http.HandlerFunc(CreateProductHandler)
	mux.Handle("/admin/product", auth.JWTMiddleware(adminOnly))
	mux.Handle("/admin/category/", auth.JWTMiddleware(http.HandlerFunc(CreateCategoryHandler)))
	mux.Handle("/categories", http.HandlerFunc(GetAllCategoriesHandler))
	mux.HandleFunc("/products", GetAllProductsHandler)
	mux.HandleFunc("/products/", GetProductByIDHandler)
	mux.HandleFunc("/products/category", GetProductByCategoryHandler)
}

package product

import (
	"fashiondemon/pkg/auth"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	adminOnly := http.HandlerFunc(CreateProductHandler)
	mux.Handle("/admin/product", auth.JWTMiddleware(adminOnly))
}

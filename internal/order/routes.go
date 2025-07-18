package order

import (
	"fashiondemon/pkg/auth"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/order", auth.JWTMiddleware(http.HandlerFunc(CreateOrderHandler)))
}

package order

import (
	"fashiondemon/pkg/auth"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	authOnly := auth.JWTMiddleware(http.HandlerFunc(CreateOrderHandler))
	mux.Handle("/orders", authOnly)

	getOrders := auth.JWTMiddleware(http.HandlerFunc(GetOrdersHandler))
	mux.Handle("/my-orders", getOrders)
}

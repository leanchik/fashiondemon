package order

import (
	"fashiondemon/pkg/auth"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	authOnly := auth.JWTMiddleware(http.HandlerFunc(CreateOrderHandler))
	mux.Handle("/orders", authOnly)
	mux.Handle("/orders/", auth.JWTMiddleware(http.HandlerFunc(GetOrderByIDHandler)))
	mux.Handle("/admin/orders/", auth.JWTMiddleware(http.HandlerFunc(GetOrderByIDAdminHandler)))

	getOrders := auth.JWTMiddleware(http.HandlerFunc(GetOrdersHandler))
	mux.Handle("/my-orders", getOrders)
}

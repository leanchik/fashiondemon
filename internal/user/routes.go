package user

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", RegisterHandler)
	mux.HandleFunc("/login", LoginHandler)
}

package user

import (
	"fashiondemon/pkg/auth"
	"fmt"
	"net/http"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(uint)
	role := r.Context().Value(auth.RoleKey).(string)

	msg := fmt.Sprintf("Ты вошел! User ID: %d | Role: %s", userID, role)
	w.Write([]byte(msg))
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", RegisterHandler)
	mux.HandleFunc("/login", LoginHandler)

	protected := http.HandlerFunc(ProtectedHandler)
	mux.Handle("/me", auth.JWTMiddleware(protected))
}

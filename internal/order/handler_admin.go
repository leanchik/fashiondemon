package order

import (
	"encoding/json"
	"fashiondemon/internal/config"
	"fashiondemon/pkg/auth"
	"net/http"
	"strconv"
	"strings"
)

func GetOrderByIDAdminHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(auth.RoleKey).(string)
	if role != "admin" {
		http.Error(w, "Доступ запрещен", http.StatusForbidden)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/admin/orders/")
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Невалидный ID заказа", http.StatusBadRequest)
		return
	}

	var order Order
	err = config.DB.Preload("Items").First(&order, orderID).Error
	if err != nil {
		http.Error(w, "Заказ не найден", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

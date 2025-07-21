package order

import (
	"encoding/json"
	"fashiondemon/internal/config"
	"fashiondemon/internal/product"
	"fashiondemon/pkg/auth"
	"net/http"
	"strconv"
	"strings"
)

type OrderInput struct {
	Item []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(uint)

	var input OrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	order := Order{
		UserID: userID,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		http.Error(w, "Ошибка при создании заказа", http.StatusInternalServerError)
		return
	}

	for _, item := range input.Item {
		var prod product.Product
		if err := config.DB.First(&prod, item.ProductID).Error; err != nil {
			http.Error(w, "Товар не найден", http.StatusNotFound)
			return
		}

		orderItem := OrderItem{
			OrderID:   order.ID,
			ProductID: prod.ID,
			Quantity:  item.Quantity,
			Price:     prod.Price,
		}
		config.DB.Create(&orderItem)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"orderID": order.ID,
		"message": "Заказ успешно оформлен",
	})
}

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(uint)

	var orders []Order
	if err := config.DB.Preload("Items").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		http.Error(w, "Ошибка при получении заказов", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func GetOrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(uint)

	idStr := strings.TrimPrefix(r.URL.Path, "/orders/")
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID заказа", http.StatusBadRequest)
		return
	}

	var order Order
	if err := config.DB.Preload("Items").Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		http.Error(w, "Заказ не найден", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

package order

import (
	"encoding/json"
	"fashiondemon/internal/config"
	"fashiondemon/internal/product"
	"fashiondemon/pkg/auth"
	"net/http"
)

type OrderInput struct {
	Item []struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
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

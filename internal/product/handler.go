package product

import (
	"encoding/json"
	"fashiondemon/internal/config"
	"fashiondemon/pkg/auth"
	"net/http"
)

type ProductInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	InStock     bool    `json:"in_stock"`
	CategoryID  uint    `json:"category_id"`
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(auth.RoleKey).(string)
	if role != "admin" {
		http.Error(w, "Доступ запрещён", http.StatusForbidden)
		return
	}

	var input ProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	product := Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		ImageURL:    input.ImageURL,
		InStock:     input.InStock,
		CategoryID:  input.CategoryID,
	}

	if result := config.DB.Create(&product); result.Error != nil {
		http.Error(w, "Ошибка при создании товара", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product

	categoryID := r.URL.Query().Get("category")
	query := config.DB

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if err := query.Find(&products).Error; err != nil {
		http.Error(w, "Ошибка при получении товаров", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

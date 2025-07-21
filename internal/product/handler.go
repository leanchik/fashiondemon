package product

import (
	"encoding/json"
	"fashiondemon/internal/config"
	"fashiondemon/pkg/auth"
	"net/http"
	"strconv"
	"strings"
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

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Невалидный ID", http.StatusBadRequest)
		return
	}

	var product Product
	if err := config.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Товар не найден", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func GetProductByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/products/category/")
	categoryID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Невалидный ID категории", http.StatusBadRequest)
		return
	}

	var products []Product
	if err := config.DB.Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		http.Error(w, "Ошибка при получении товаров", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(auth.RoleKey).(string)
	if role != "admin" {
		http.Error(w, "Доступ запрещен!", http.StatusForbidden)
	}

	idStr := strings.TrimPrefix(r.URL.Path, "admin/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	if err := config.DB.Delete(&Product{}, id).Error; err != nil {
		http.Error(w, "Ошибка при удалении товара", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Товар удален!"))
}

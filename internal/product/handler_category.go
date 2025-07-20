package product

import (
	"encoding/json"
	"fashiondemon/internal/config"
	"fashiondemon/pkg/auth"
	"net/http"
)

type CategoryInput struct {
	Name string `json:"name"`
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(auth.RoleKey).(string)
	if role != "admin" {
		http.Error(w, "Доступ запрещен", http.StatusForbidden)
		return
	}

	var input CategoryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	category := Category{Name: input.Name}
	if err := config.DB.Create(&category).Error; err != nil {
		http.Error(w, "Ошибка при создании категории", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

func GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var categories []Category

	if err := config.DB.Find(&categories).Error; err != nil {
		http.Error(w, "Ошибка при получении категорий", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

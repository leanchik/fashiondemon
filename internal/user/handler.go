package user

import (
	"encoding/json"
	"fashiondemon/internal/config"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var input RegisterInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Ошибка хеширования пароля", http.StatusInternalServerError)
		return
	}

	user := User{
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		Role:         "user",
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "Не удалось создать пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Пользователь зарегистрирован"))
}

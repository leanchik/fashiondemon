package config

import (
	"fashiondemon/internal/order"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе", err)
	}

	DB = db

	err = db.AutoMigrate(
		&order.Order{},
		&order.OrderItem{},
	)
	if err != nil {
		log.Fatal("Ошибка миграции", err)
	}

	fmt.Println("Успешное подключение к базе и миграция прошла!")

}

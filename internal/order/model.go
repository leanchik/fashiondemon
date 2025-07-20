package order

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	Total      float64
	Status     string `gorm:"default:'new'"`
	CreateAt   time.Time
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
}

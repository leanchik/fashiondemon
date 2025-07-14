package order

import "time"

type Order struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	Total      float64
	Status     string `gorm:"default:'new'"`
	CreateAt   time.Time
	OrderItems []OrderItem
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  uint
	UnitPrice float64
}

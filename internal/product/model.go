package product

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	ImageURL    string
	InStock     bool
	CategoryID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

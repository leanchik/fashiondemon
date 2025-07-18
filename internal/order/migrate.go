package order

import "fashiondemon/internal/config"

func Migrate() {
	config.DB.AutoMigrate(&Order{}, &OrderItem{})
}

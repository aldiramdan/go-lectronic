package models

import "time"

type Cart struct {
	CartID      string     `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty"`
	UserID      string     `gorm:"type:uuid" json:"user_id"`
	ProductID   string     `gorm:"type:uuid" json:"product_id"`
	Status      string     `gorm:"type:varchar(255)" json:"status"`
	TotalPrice  int64      `gorm:"type:int" json:"total_price"`
	CreatedAt   time.Time  `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamp" json:"updated_at"`
}

type Carts []Cart

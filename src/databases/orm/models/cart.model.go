package models

import "time"

type Cart struct {
	CartID    string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty"`
	UserID    string    `gorm:"type:uuid" json:"user_id"`
	User      User      `json:"user"`
	ProductID string    `gorm:"type:uuid" json:"product_id"`
	Product   Product   `json:"product"`
	Discount  int64     `gorm:"type:int" json:"discount,omitempty"`
	Amount    int64     `gorm:"type:int" json:"amount,omitempty"`
	Total     int64     `gorm:"type:int" json:"total"`
	Status    string    `gorm:"type:varchar(255)" json:"status"`
	Payment   string    `gorm:"type:varchar(255)" json:"payment,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

type Carts []Cart

func (Cart) TableName() string {
	return "carts"
}

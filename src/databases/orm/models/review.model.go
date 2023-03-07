package models

import "time"

type Review struct {
	ReviewID  string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty"`
	UserID    string    `gorm:"type:uuid" json:"user_id"`
	User      User      `json:"user"`
	ProductID string    `gorm:"type:uuid" json:"product_id"`
	Product   Product   `json:"product,omitempty"`
	Comment   string    `gorm:"type:text" json:"comment"`
	Rating    float64   `gorm:"type:numeric" json:"rating"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

type Reviews []Review

func (Review) TableName() string {
	return "reviews"
}

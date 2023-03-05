package models

import (
	"time"
)

type User struct {
	UserID       string    `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Username     string    `json:"username,omitempty" valid:"type(string),required~Username is needed"`
	Email        string    `json:"email" valid:"email,required~Email is needed"`
	Password     string    `json:"password,omitempty"`
	Gender       string    `json:"gender,omitempty" valid:"-"`
	Address      string    `json:"address,omitempty" valid:"-"`
	DateOfBirth  string    `json:"date_of_birth,omitempty" valid:"-"`
	Role         string    `gorm:"default:user" json:"role,omitempty" valid:"-"`
	Image        string    `json:"image,omitempty" valid:"-"`
	MobileNumber string    `json:"mobile_number,omitempty" valid:"-"`
	TokenVerify  string    `json:"token_verify,omitempty" valid:"-"`
	IsActive     bool      `gorm:"default: false" json:"is_active,omitempty" valid:"-"`
	CreatedAt    time.Time `json:"created_at"  valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

type Users []User

func (User) TableName() string {
	return "users"
}

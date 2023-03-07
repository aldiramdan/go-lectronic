package interfaces

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
)

type CartRepo interface {
	GetAll() (*models.Carts, error)
	GetByID(id string) (*models.Cart, error)
	GetByUserID(id string) (*models.Carts, error)
	Add(cart *models.Cart) (*models.Cart, error)
	Update(id string, cart *models.Cart) (*models.Cart, error)
	Delete(id string) (*models.Cart, error)
}

type CartSrvc interface {
	GetAll() *libs.Response
	GetByID(id string) *libs.Response
	GetByUserID(id string) *libs.Response
	Add(cart *models.Cart) *libs.Response
	Update(id string, cart *models.Cart) *libs.Response
	Delete(id string) *libs.Response
}

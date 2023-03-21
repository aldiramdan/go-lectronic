package interfaces

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
)

type ProductRepo interface {
	GetAll() (*models.Products, error)
	GetByID(id string) (*models.Product, error)
	Popular() (*models.Products, error)
	Search(query string) (*models.Products, error)
	Add(product *models.Product) (*models.Product, error)
	Update(id string, product *models.Product) (*models.Product, error)
	Delete(id string) (*models.Product, error)
}

type ProductSrvc interface {
	GetAll() *libs.Response
	GetByID(id string) *libs.Response
	Popular() *libs.Response
	Search(query string) *libs.Response
	Add(product *models.Product) *libs.Response
	Update(id string, product *models.Product) *libs.Response
	Delete(id string) *libs.Response
}

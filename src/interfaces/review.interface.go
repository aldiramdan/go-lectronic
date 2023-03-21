package interfaces

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
)

type ReviewRepo interface {
	GetByID(id string) (*models.Review, error)
	GetByProductID(id string) (*models.Reviews, error)
	Add(review *models.Review) (*models.Review, error)
}

type ReviewService interface {
	GetByID(id string) *libs.Response
	GetByProductID(id string) *libs.Response
	Add(review *models.Review) *libs.Response
}

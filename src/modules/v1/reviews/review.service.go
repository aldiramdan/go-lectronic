package reviews

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/interfaces"
	"lectronic/src/libs"

	"gorm.io/gorm"
)

type ReviewService struct {
	repo interfaces.ReviewRepo
}

func NewService(repo interfaces.ReviewRepo) *ReviewService {
	return &ReviewService{repo}
}

func (s *ReviewService) GetByID(id string) *libs.Response {
	result, err := s.repo.GetByID(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	return libs.GetResponse(result, 200, false)
}

func (s *ReviewService) GetByProductID(id string) *libs.Response {
	result, err := s.repo.GetByProductID(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *ReviewService) Add(review *models.Review) *libs.Response {
	result, err := s.repo.Add(review)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

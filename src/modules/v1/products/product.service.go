package products

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/interfaces"
	"lectronic/src/libs"

	"gorm.io/gorm"
)

type ProductService struct {
	repo interfaces.ProductRepo
}

func NewService(repo interfaces.ProductRepo) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) GetAll() *libs.Response {
	result, err := s.repo.GetAll()

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *ProductService) GetByID(id string) *libs.Response {
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

func (s *ProductService) Add(product *models.Product) *libs.Response {
	result, err := s.repo.Add(product)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *ProductService) Update(id string, product *models.Product) *libs.Response {

	_, err := s.repo.GetByID(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	result, err := s.repo.Update(id, product)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *ProductService) Delete(id string) *libs.Response {

	_, err := s.repo.GetByID(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	_, err = s.repo.Delete(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{"message": "product deleted successfully"}

	return libs.GetResponse(response, 200, false)
}

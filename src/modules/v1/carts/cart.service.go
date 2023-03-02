package carts

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/interfaces"
	"lectronic/src/libs"

	"gorm.io/gorm"
)

type CartService struct {
	repo interfaces.CartRepo
}

func NewService(repo interfaces.CartRepo) *CartService {
	return &CartService{repo}
}

func (s *CartService) GetAll() *libs.Response {
	result, err := s.repo.GetAll()

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *CartService) GetByID(id string) *libs.Response {
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

func (s *CartService) GetByUserID(id string) *libs.Response {
	result, err := s.repo.GetByUserID(id)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *CartService) Add(cart *models.Cart) *libs.Response {

	result, err := s.repo.Add(cart)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *CartService) Update(id string, cart *models.Cart) *libs.Response {

	_, err := s.repo.GetByID(id)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return libs.GetResponse(err.Error(), 404, true)
		default:
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	result, err := s.repo.Update(id, cart)

	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *CartService) Delete(id string) *libs.Response {

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

	response := map[string]string{"message": "cart deleted successfully"}

	return libs.GetResponse(response, 200, false)
}

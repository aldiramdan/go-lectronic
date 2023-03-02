package carts

import (
	"errors"
	"lectronic/src/databases/orm/models"

	"gorm.io/gorm"
)

type CartRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *CartRepo {
	return &CartRepo{db}
}

func (r *CartRepo) GetAll() (*models.Carts, error) {
	var carts models.Carts

	if err := r.db.
		Order("created_at DESC").
		Find(&carts).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(carts) <= 0 {
		return nil, errors.New("data cart is empty")
	}

	return &carts, nil
}

func (r *CartRepo) GetByID(id string) (*models.Cart, error) {
	var cart models.Cart

	if err := r.db.
		First(&cart, "cart_id = ?", id).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *CartRepo) GetByUserID(id string) (*models.Cart, error) {
	var cart models.Cart

	if err := r.db.
		Where("user_id = ?", id).Find(&cart).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *CartRepo) Add(cart *models.Cart) (*models.Cart, error) {
	if err := r.db.
		Create(cart).
		Find(&cart).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	return cart, nil
}

func (r *CartRepo) Update(id string, cart *models.Cart) (*models.Cart, error) {
	if err := r.db.
		Model(&cart).
		Where("cart_id = ?", id).
		Updates(&cart).
		Find(&cart).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return cart, nil
}

func (r *CartRepo) Delete(id string) (*models.Cart, error) {
	var cart models.Cart

	if err := r.db.
		Delete(cart, "cart_id = ?", id).Error; err != nil {
		return nil, errors.New("failed to delete data")
	}

	return &cart, nil
}

package products

import (
	"errors"
	"lectronic/src/databases/orm/models"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (r *ProductRepo) GetAll() (*models.Products, error) {
	var products models.Products

	if err := r.db.
		Order("created_at DESC").
		Find(&products).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	if len(products) <= 0 {
		return nil, errors.New("data product is empty")
	}

	return &products, nil
}

func (r *ProductRepo) GetByID(id string) (*models.Product, error) {
	var product models.Product

	if err := r.db.
		First(&product, "product_id = ?", id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepo) Add(product *models.Product) (*models.Product, error) {
	if err := r.db.
		Create(product).
		Find(&product).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	return product, nil
}

func (r *ProductRepo) Update(id string, product *models.Product) (*models.Product, error) {
	if err := r.db.
		Model(&product).
		Where("product_id = ?", id).
		Updates(&product).
		Find(&product).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return product, nil
}

func (r *ProductRepo) Delete(id string) (*models.Product, error) {
	var product models.Product

	if err := r.db.
		Delete(product, "product_id = ?", id).Error; err != nil {
		return nil, errors.New("failed to delete data")
	}

	return &product, nil
}

package reviews

import (
	"errors"
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"

	"gorm.io/gorm"
)

type ReviewRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *ReviewRepo {
	return &ReviewRepo{db}
}

func (r *ReviewRepo) GetByID(id string) (*models.Review, error) {
	var review models.Review

	if err := r.db.
		Find(&review, "review_id = ?", id).Error; err != nil {
		return nil, err
	}

	return &review, nil
}

func (r *ReviewRepo) GetByProductID(id string) (*models.Reviews, error) {
	var reviews models.Reviews

	if err := r.db.
		Order("created_at DESC").
		Where("product_id = ?", id).
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	if len(reviews) <= 0 {
		return nil, errors.New("data cart is empty")
	}

	return &reviews, nil
}

func (r *ReviewRepo) Add(review *models.Review) (*models.Review, error) {

	var product models.Product
	var count int64

	if err := r.db.
		First(&product, "product_id = ?", review.ProductID).Error; err != nil {
		return nil, errors.New("data product not found")
	}

	if err := r.db.
		Model(&review).Where("product_id = ?", review.ProductID).Count(&count).Error; err != nil {
		return nil, errors.New("data product not found")
	}

	count = count + 1

	rating := libs.CalculateNewRating(count, product.Rating, review.Rating)

	if err := r.db.
		Create(review).
		Find(&review).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	if err := r.db.
		Model(&product).
		Where("product_id = ?", review.ProductID).
		Update("rating", rating).Error; err != nil {
		return nil, errors.New("failed to update data rating")
	}

	return review, nil
}

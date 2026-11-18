package repository

import (
	"go-project/internal/domain"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product

	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product

	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}
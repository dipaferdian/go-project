package usecase

import (
	"errors"

	"go-project/internal/domain"
)

var (
	ErrProductNameRequired = errors.New("product name is required")
	ErrInvalidProductPrice = errors.New("product price must be greater than zero")
)

type ProductUsecase struct {
	repository domain.ProductRepository
}

func NewProductUsecase(repository domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		repository: repository,
	}
}

func (u *ProductUsecase) GetProducts() ([]domain.Product, error) {
	return u.repository.FindAll()
}

func (u *ProductUsecase) GetProductByID(id uint) (*domain.Product, error) {
	return u.repository.FindByID(id)
}

func (u *ProductUsecase) CreateProduct(product *domain.Product) error {
	if product.Name == "" {
		return ErrProductNameRequired
	}

	if product.Price <= 0 {
		return ErrInvalidProductPrice
	}

	return u.repository.Create(product)
}
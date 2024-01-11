package service

import (
	"Ejercicio_1/internal/domain"
	"Ejercicio_1/internal/repository"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{productRepository}
}

func (s *ProductServiceImpl) AddProduct(product *domain.Product) (*domain.Product, error) {
	// Validar la fecha de vencimiento
	return s.productRepository.AddProduct(product)
}

func (s *ProductServiceImpl) GetProducts() []domain.Product {
	return s.productRepository.GetProducts()
}

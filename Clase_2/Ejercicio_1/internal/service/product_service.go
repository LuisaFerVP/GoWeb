package service

import "Ejercicio_1/internal/domain"

type ProductService interface {
	AddProduct(product *domain.Product) (*domain.Product, error)
	GetProducts() []domain.Product
}

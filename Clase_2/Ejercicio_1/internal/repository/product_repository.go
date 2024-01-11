package repository

import "Ejercicio_1/internal/domain"

type ProductRepository interface {
	AddProduct(product *domain.Product) (*domain.Product, error)
	GetProducts() []domain.Product
}

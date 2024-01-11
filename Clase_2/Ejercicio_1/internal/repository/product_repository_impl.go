package repository

import (
	"Ejercicio_1/internal/domain"
)

type ProductRepositoryImpl struct {
	products []domain.Product
}

func NewProductRepository(products []domain.Product) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{products}
}

func (r *ProductRepositoryImpl) AddProduct(product *domain.Product) (*domain.Product, error) {
	// Generar un nuevo ID para el producto
	product.ID = generateNewID(r.products)
	r.products = append(r.products, *product)
	return product, nil
}

func (r *ProductRepositoryImpl) GetProducts() []domain.Product {
	return r.products
}

func generateNewID(products []domain.Product) int {
	highestID := 0
	for _, p := range products {
		if p.ID > highestID {
			highestID = p.ID
		}
	}
	return highestID + 1
}

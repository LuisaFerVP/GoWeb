package handler

import (
	"encoding/json"
	"net/http"

	"Ejercicio_1/internal/domain"
	"Ejercicio_1/internal/service"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct domain.Product

	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	addedProduct, err := h.productService.AddProduct(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error adding product"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addedProduct)
}

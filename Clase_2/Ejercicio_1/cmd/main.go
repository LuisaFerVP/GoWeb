package main

import (
	"fmt"
	"log"
	"net/http"

	"Ejercicio_1/internal/handler"
	"Ejercicio_1/internal/repository"
	"Ejercicio_1/internal/service"
	"Ejercicio_1/internal/utils"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Cargar productos desde JSON
	products, err := utils.LoadProductsFromJSON("data/products.json")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Configurar el enrutador
	router := chi.NewRouter()

	// Inyectar dependencias
	productRepository := repository.NewProductRepository(products)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	// Rutas
	router.Post("/products", productHandler.AddProduct)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Products struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_Value   string  `json:"code_value"`
	Is_Published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

func main() {
	// Lee el archivo JSON
	data, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Decodifica el JSON en una slice de productos
	var products []Products
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Configurar el enrutador
	router := chi.NewRouter()

	// Ruta /ping
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// Ruta /products
	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		// Devolver la lista de todos los productos en formato JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})

	// Ruta /products/:id
	router.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Obtener el parámetro de la URL
		idParam := chi.URLParam(r, "id")

		// Convertir el parámetro a un entero
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid ID"))
			return
		}

		// Buscar el producto por ID
		var foundProduct *Products
		for _, p := range products {
			if p.ID == id {
				foundProduct = &p
				break
			}
		}

		// Responder con el producto encontrado o un mensaje de error si no se encuentra
		if foundProduct != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(foundProduct)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Product not found"))
		}
	})

	// Ruta /products/search
	router.Get("/products/search", func(w http.ResponseWriter, r *http.Request) {
		// Obtener el parámetro de la URL
		priceGtParam := r.URL.Query().Get("priceGt")

		// Convertir el parámetro a un número flotante
		priceGt, err := strconv.ParseFloat(priceGtParam, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid priceGt parameter"))
			return
		}

		// Filtrar productos por precio mayor al valor proporcionado
		var filteredProducts []Products
		for _, p := range products {
			if p.Price > priceGt {
				filteredProducts = append(filteredProducts, p)
			}
		}

		// Responder con la lista de productos filtrados
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filteredProducts)
	})

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

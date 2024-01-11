package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	// Ruta /products (POST)
	router.Post("/products", func(w http.ResponseWriter, r *http.Request) {
		// Decodificar la solicitud JSON
		var newProduct Products
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error decoding JSON"))
			return
		}

		// Validar los campos del nuevo producto
		if newProduct.Name == "" || newProduct.Code_Value == "" || newProduct.Expiration == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Name, Code_Value, and Expiration cannot be empty"))
			return
		}

		// Verificar si el Code_Value ya existe en la lista
		for _, p := range products {
			if p.Code_Value == newProduct.Code_Value {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Code_Value must be unique"))
				return
			}
		}

		// Generar un nuevo ID para el producto
		newProduct.ID = generateNewID(products)

		// Validar la fecha de vencimiento
		_, err = time.Parse("02/01/2006", newProduct.Expiration)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid expiration date format"))
			return
		}

		// Agregar el nuevo producto a la lista
		products = append(products, newProduct)

		// Responder con el nuevo producto en formato JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newProduct)
	})

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error:", err)
	}
}

// Función auxiliar para generar un nuevo ID único
func generateNewID(products []Products) int {
	highestID := 0
	for _, p := range products {
		if p.ID > highestID {
			highestID = p.ID
		}
	}
	return highestID + 1
}

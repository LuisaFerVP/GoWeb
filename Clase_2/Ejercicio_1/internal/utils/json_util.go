package utils

import (
	"Ejercicio_1/internal/domain"
	"encoding/json"
	"io/ioutil"
	"log"
)

// LoadProductsFromJSON carga productos desde un archivo JSON
func LoadProductsFromJSON(filePath string) ([]domain.Product, error) {
	// Leemos el archivo desde la ruta proporcionada
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error leyendo el archivo JSON: %v", err)
		return nil, err
	}

	// Parseamos los datos JSON y devolvemos la lista de productos
	var products []domain.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Printf("Error decodificando el JSON: %v", err)
		return nil, err
	}

	return products, nil
}

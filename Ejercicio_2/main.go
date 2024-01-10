package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// GreetingRequest estructura para la solicitud de saludo
type GreetingRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	// router
	router := chi.NewRouter()

	// Endpoint /ping
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		// request
		fmt.Println("GET /ping")
		fmt.Println("method:", r.Method)
		fmt.Println("url:", r.URL)
		fmt.Println("header:", r.Header)

		w.Write([]byte(`{"message": "pong"}`))
	})

	// Endpoint /greetings
	router.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		// Decodificar la solicitud JSON
		var greetingReq GreetingRequest
		err := json.NewDecoder(r.Body).Decode(&greetingReq)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		// Construir el saludo
		greeting := fmt.Sprintf("Hello %s %s", greetingReq.FirstName, greetingReq.LastName)

		// Responder con el saludo en formato JSON
		response := map[string]string{"greeting": greeting}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

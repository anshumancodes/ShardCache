package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
}

func HelloFrom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	response := &Response{Message: "http server is running sucessfully!"}

	json.NewEncoder(w).Encode(response)

}
func main() {
	// loading from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// initialize the server
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/", HelloFrom)

	// port and server running configuration
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	fmt.Printf("server is running on %s\n", PORT)
	serverError := http.ListenAndServe(":"+PORT, mux)
	if serverError != nil {
		fmt.Println(serverError)
	}

}

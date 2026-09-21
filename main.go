package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/anshumancodes/ShardCache/cache"
	cachehttp "github.com/anshumancodes/ShardCache/http"
	"github.com/joho/godotenv"
)

func main() {
	// loading from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	c := cache.New(6)

	handler := cachehttp.NewHandler(c)

	// initialize the server
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("GET /", cachehttp.HelloFrom)
	// cache routes
	mux.HandleFunc("GET /cache/{key}", handler.Get)
	mux.HandleFunc("POST /cache/{key}", handler.Set)

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

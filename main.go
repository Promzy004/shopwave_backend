package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/Promzy004/shopwave_backend.git/internal/routes"
)

func main () {
	r := chi.NewRouter()
	routes.InitRoutes(r)

	fmt.Println("Running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
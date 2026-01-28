package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Promzy004/shopwave_backend.git/config"
	"github.com/Promzy004/shopwave_backend.git/internal/routes"
	"github.com/go-chi/chi/v5"
)

func main () {
	config.LoadEnv()

	port := os.Getenv("PORT")
	r := chi.NewRouter()
	routes.InitRoutes(r)

	config.ConnectDB()
	defer config.DB.Close()

	fmt.Println("Server is running ...")
	fmt.Printf("Running server on [http://127.0.0.1:%s]\n", port)
	log.Fatal(http.ListenAndServe(":8000", r))
}
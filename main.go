package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Promzy004/shopwave_backend.git/config"
	"github.com/Promzy004/shopwave_backend.git/internal/routes"
	"github.com/go-chi/chi/v5"
	// "github.com/google/uuid"
)

func main () {
	config.LoadEnv()
	
	migrateFresh := flag.Bool("fresh", false, "Drop all tables and recreate them")
    flag.Parse()

	port := os.Getenv("PORT")
	r := chi.NewRouter()
	routes.InitRoutes(r)

	config.ConnectDB()
	defer config.DB.Close()

	if *migrateFresh {
        config.FreshMigrate()
    } else {
        config.InitDB()
    }

	fmt.Println("Server is running ...")
	fmt.Printf("Running server on [http://127.0.0.1:%s]\n", port)
	log.Fatal(http.ListenAndServe(":8000", r))
}
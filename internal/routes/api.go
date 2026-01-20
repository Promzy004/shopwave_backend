package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(r *chi.Mux) {
	r.Get("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world, promise to add more routes soon!"))
	})
}
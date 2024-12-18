package routes

import (
	"aquamaster/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	r.HandleFunc("/profile", handlers.ProfileHandler).Methods("GET")
	r.HandleFunc("/cart", handlers.CartHandler).Methods("GET")
	r.HandleFunc("/favorites", handlers.FavoritesHandler).Methods("GET")

	r.HandleFunc("/plumbing", handlers.PlumbingHandler).Methods("GET")
	r.HandleFunc("/interior", handlers.InteriorHandler).Methods("GET")
	r.HandleFunc("/accessories", handlers.AccessoriesHandler).Methods("GET")

	return r
}

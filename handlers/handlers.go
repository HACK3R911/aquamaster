package handlers

import (
	"html/template"
	"net/http"
)

// /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ui/html/pages/MainMenu.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// /profile
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ui/html/pages/Profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// /cart
func CartHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/html/pages/cart.html")
}

// /favorites
func FavoritesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/html/pages/favorites.html") // Страница избранного
}

// /accessories
func AccessoriesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/html/pages/accessories.html")
}

func InteriorHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/html/pages/interior.html")
}

func PlumbingHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/html/pages/plumbing.html")
}

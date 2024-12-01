package main

import (
	"net/http"
)

// setupRoutes устанавливает маршруты для приложения
func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/html/pages/MainMenu.html") // Главная страница
	})

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/html/pages/Profile.html") // Профиль пользователя
	})

	http.HandleFunc("/favorites", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/html/pages/favorites.html") // Страница избранного
	})

	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/html/pages/cart.html") // Страница корзины
	})

	http.HandleFunc("/accessories", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/html/pages/accessories.html") // Страница аксессуаров
	})

	http.HandleFunc("/interior", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/html/pages/interior.html") // Страница товаров для интерьера
	})

	http.HandleFunc("/plumbing", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/html/pages/plumbing.html") // Страница сантехники
	})
}

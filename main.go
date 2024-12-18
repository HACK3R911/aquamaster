package main

import (
	"aquamaster/routes"
	"log"
	"net/http"
)

func main() {
	//static files
	fs := http.FileServer(http.Dir("./ui/html/pages/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Настройка маршрутов
	r := routes.SetupRoutes()

	// Запуск сервера на порту 8080
	log.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"
)

func main() {
	// Указываем путь к статическим файлам (CSS, JS и т.д.)
	fs := http.FileServer(http.Dir("./ui/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Настройка маршрутов
	setupRoutes()

	// Запуск сервера на порту 8080
	log.Println("Сервер запущен на :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

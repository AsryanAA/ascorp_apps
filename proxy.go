package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main23() {
	// Настройка прокси
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "example.com", // Замените на нужный хост или оставьте пустым для перехвата всех запросов
	})

	// Обработчик для логирования запросов
	proxy.ModifyResponse = func(resp *http.Response) error {
		log.Printf("Response: %s", resp.Request.URL)
		return nil
	}

	// Обработчик для логирования запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL)
		proxy.ServeHTTP(w, r)
	})

	// Запуск сервера на порту 8080
	log.Println("Proxy server running on :8080")
	log.Fatal(http.ListenAndServe(":8016", nil))
}
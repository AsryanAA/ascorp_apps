package main

import (
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Статические файлы (CSS, JS, изображения)
	router.Static("/static", "./static")

	router.GET("/gidrant", func(c *gin.Context) {
		// Проверяем, существует ли запрашиваемый файл
		filePath := path.Join("./static", c.Request.URL.Path)
		
		// Если файл существует и это не API запрос - отдаем его
		if c.Request.URL.Path != "/" && !isAPIRequest(c) && fileExists(filePath) {
			c.File(filePath)
			return
		}
		
		// Для всех остальных случаев отдаем index.html
		c.File("./static/index.html")
	})

	// API endpoint пример
	router.GET("/api/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time":    "2024-01-15 10:30:00",
			"message": "Данные из API",
		})
	})

	log.Println("Server starting on :8080")
	router.Run(":8080")
}

// Проверяем, является ли запрос API запросом
func isAPIRequest(c *gin.Context) bool {
	return path.Base(c.Request.URL.Path) == "api" || 
	       path.Dir(c.Request.URL.Path) == "/api"
}

// Проверяем существование файла
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
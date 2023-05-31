package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammadali07/go_project/api"
)

func main() {
	// Inisialisasi router Gin
	router := gin.Default()

	// Setup router API
	apiRouter := api.SetupRouter()

	// Tambahkan router API ke router utama
	router.Use(apiRouter.ServeHTTP)

	// Jalankan server pada port 8080
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

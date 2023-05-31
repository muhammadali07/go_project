package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadali07/go_project/app"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Definisikan rute-rute API
	router.GET("/users", app.GetUsers)
	router.GET("/users/:id", app.GetUser)
	router.POST("/users", app.CreateUser)
	router.PUT("/users/:id", app.UpdateUser)
	router.DELETE("/users/:id", app.DeleteUser)

	return router
}

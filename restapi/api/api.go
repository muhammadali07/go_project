package api

import (
	"restapi/app"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Definisikan rute-rute API
	router.GET("/users", app.GetUsers)
	router.GET("/users/:id", app.GetUserID)
	router.POST("/users", app.CreateUser)
	router.PUT("/users/:id", app.UpdateUserID)
	router.DELETE("/users/:id", app.DeleteUser)

	return router
}

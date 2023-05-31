package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadali07/go_project/repository"
)

// Handler untuk mendapatkan semua pengguna
func getUsers(c *gin.Context) {
	var users []repository.User
	err := repository.GetUsers(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, repository.APIResponse{Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, repository.APIResponse{Success: true, Data: users})
}

// Handler untuk mendapatkan pengguna berdasarkan ID
func getUser(c *gin.Context) {
	id := c.Param("id")
	user, err := repository.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, repository.APIResponse{Success: false, Error: "User not found"})
		return
	}
	c.JSON(http.StatusOK, repository.APIResponse{Success: true, Data: user})
}

// Handler untuk membuat pengguna baru
func createUser(c *gin.Context) {
	var user repository.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, repository.APIResponse{Success: false, Error: err.Error()})
		return
	}
	err := repository.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, repository.APIResponse{Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, repository.APIResponse{Success: true, Data: user})
}

// Handler untuk memperbarui pengguna
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user repository.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, repository.APIResponse{Success: false, Error: "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.UpdateUser(id, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, repository.APIResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, repository.APIResponse{Success: true, Data: user})
}

// Handler untuk menghapus pengguna
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := repository.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, repository.APIResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, repository.APIResponse{Success: true, Message: "User deleted successfully"})
}

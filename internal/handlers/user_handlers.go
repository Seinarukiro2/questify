package handlers

import (
	"net/http"
	"questify/internal/functions"

	"github.com/gin-gonic/gin"
)

func CreateNewUserHandler(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	if name == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметры 'name' и 'password' обязательны"})
		return
	}

	err := functions.CreateUserAndPrint(name, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь был создан"})
}

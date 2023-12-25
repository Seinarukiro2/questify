package handlers

import (
	"net/http"
	"questify/internal/functions"

	"github.com/gin-gonic/gin"
)

func CreateNewUserHandler(c *gin.Context) {
	name := "Timur"
	password := "123"
	err := functions.CreateUserAndPrint(name, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User was created"})
}

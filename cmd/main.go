package main

import (
	// "net/http"

	handler "questify/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingPongHandler)
	r.GET("/create", handler.CreateNewUserHandler)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

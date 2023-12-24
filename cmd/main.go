package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"questify/internal/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingPongHandler)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

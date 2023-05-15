package main

import (
	"example.com/jwt/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.GET("/api/rest/Auth/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data":    nil,
			"message": "pong",
			"status":  http.StatusOK,
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}

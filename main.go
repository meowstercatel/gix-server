package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	auth := r.Group("/auth")
	auth.POST("/login", func(ctx *gin.Context) {})
	auth.POST("/register", func(ctx *gin.Context) {})

	repo := r.Group("/repo/:id")
	repo.GET("/", func(ctx *gin.Context) {})
	repo.POST("/push", func(ctx *gin.Context) {})

	r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simonha9/rate-limiter/pkg/middleware"

)

func main() {
	r := gin.Default()

	tokenLimiter := middleware.NewTokenLimiter()

	r.Use(tokenLimiter.TokenLimiterMiddleware())
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	go tokenLimiter.Run()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
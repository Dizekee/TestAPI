package main

import (
	"net/http"

	"github.com/Dizekee/TestAPI/src/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitEnv()
	config.Connection()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Messsage": "HELLO there boys",
		})
	})

	router.Run(":8080")

	router.Run()
}

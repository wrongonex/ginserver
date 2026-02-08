package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Messsage": "Are you still doubtful about the power of Go?, babes that's cold",
		})
	})

	router.Run(":3000")

	router.Run()
}

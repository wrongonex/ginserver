package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Messsage": "I'm running all fine on port 3000",
		})
	})

	router.Run(":3000")

	router.Run()
}

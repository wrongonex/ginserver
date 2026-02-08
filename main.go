package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Messsage": "WOooohhaaaaooo, this is gonna be a magical journey",
		})
	})

	router.Run(":3000")

	router.Run()
}

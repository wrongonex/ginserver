package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Messsage": "I have no limitations",
	})
}

func RandomMsg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Messsage": "Random Message",
	})
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Messsage": "This is not overkill, this is gin server",
	})
}

func RandomMsg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Messsage": "Random Message",
	})
}

func TestRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success":"true",
		"Hello":"gta 6 is coming soon...",
	})
}

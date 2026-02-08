package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wrongonex/ginserver/routes"
)

func main() {
	router := gin.Default()

	router.GET("/", routes.MainHandler)
	router.GET("/random", routes.RandomMsg)
	router.GET("/test", routes.TestRoute)

	router.Run(":3000")

	router.Run()
}

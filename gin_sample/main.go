package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qjw/middleware"
)

func main() {
	r := gin.Default()
	r.GET("/nocache", middleware.GinHandle(middleware.NoCache),
		func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:9090") // listen and serve on
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jsonw23/copilot-react-admin/middleware"
)

func main() {
	r := gin.Default()

	// use the auth middleware
	r.Use(middleware.Auth(r))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World",
		})
	})

	r.Run()
}

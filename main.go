package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jsonw23/copilot-react-admin/middleware"
)

func main() {
	r := gin.Default()

	// use the auth middleware
	middleware.InitAuthStore(r)
	r.Use(middleware.Auth())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World",
		})
	})

	r.Run()
}

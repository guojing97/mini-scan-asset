package main

import "github.com/gin-gonic/gin"

func main() {
	routes := gin.Default()

	routes.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	routes.Run()
}

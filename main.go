package main

import "github.com/gin-gonic/gin"

func main() {
	// Create a Gin router
	router := gin.Default()

	// Define a GET endpoint
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}

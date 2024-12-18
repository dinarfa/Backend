package main

import (
	"final-project/database"
	"final-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Halo!",
		})
	})

	// apiKey := "452596ab9d4425c50c2b265a0a88af9e"
	db := database.ConnectToDb()
	routes.Routes(r, db)
	r.Run(":8080")
}

package main

import (
	"log"
	"os"

	"github.com/Truong62/taskoria/auth-service/config"
	"github.com/Truong62/taskoria/auth-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	config.ConnectDB()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Auth Server s")
	})

	routes.AuthRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server started at http://localhost:" + port)
	r.Run(":" + port)
}

package main

import (
	"fmt"
	"go-docker-devcontainer-starter/middleware"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8000"

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Unable To Load .env File")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()
	r.Use(middleware.GinContextToContextMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	fmt.Println("App running on port: " + defaultPort)
}

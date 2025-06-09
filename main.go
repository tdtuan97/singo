package main

import (
	"os"
	"singo/conf"
	"singo/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from config file
	conf.Init()

	// Load routes
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := server.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

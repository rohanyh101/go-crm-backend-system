package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/roh4nyh/matrice_ai/routes"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Printf("error loading .env file: %v", err)
	// }

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	gin.SetMode(gin.ReleaseMode)

	app := gin.New()
	app.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	app.Use(cors.New(config))

	routes.AuthRoutes(app)

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "server is up and running..."})
	})

	routes.UserRoutes(app)

	routes.CustomerRoutes(app)

	app.Run(fmt.Sprintf(":%s", PORT))
}

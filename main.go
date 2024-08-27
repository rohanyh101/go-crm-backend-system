package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	// _ "github.com/roh4nyh/matrice_ai/docs"
	"github.com/roh4nyh/matrice_ai/routes"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
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

	router := gin.New()
	router.Use(gin.Logger())

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.AuthRoutes(router)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "server is up and running..."})
	})

	routes.UserRoutes(router)

	routes.CustomerRoutes(router)

	router.Run(fmt.Sprintf(":%s", PORT))
}

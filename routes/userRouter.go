package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rohanhonnakatti/golang-jwt-auth/controllers"
	"github.com/rohanhonnakatti/golang-jwt-auth/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.AuthenticateUser())

	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.PUT("/users/:user_id", controller.UpdateUser())
	incomingRoutes.DELETE("/users/:user_id", controller.DeleteUser())

	incomingRoutes.POST("/users/meet/", controller.CreateInteraction())
	incomingRoutes.DELETE("/users/meet/:meet_id", controller.DeleteInteraction())
	incomingRoutes.GET("/users/meet/", controller.GetInteractionsByUserID())
}

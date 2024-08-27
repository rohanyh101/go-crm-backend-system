package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/roh4nyh/matrice_ai/controllers"
	"github.com/roh4nyh/matrice_ai/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.AuthenticateUser())

	// user crud operations
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.PUT("/users/:user_id", controller.UpdateUser())
	incomingRoutes.DELETE("/users/:user_id", controller.DeleteUser())

	// user services
	// get all interactions, only for admin
	incomingRoutes.GET("/users/meet/", controller.GetAllInteractions())

	// get all interactions by user id
	incomingRoutes.GET("/user/meet/", controller.GetInteractionsByUserID())

	// get all interactions by customer id
	incomingRoutes.POST("/users/meet/:customer_id", controller.CreateInteractionAndSendEmail())

	// delete interaction by meet id
	incomingRoutes.DELETE("/users/meet/:interaction_id", controller.DeleteInteraction())
}

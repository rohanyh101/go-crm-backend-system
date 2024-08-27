package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/roh4nyh/matrice_ai/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	// user authentication
	incomingRoutes.POST("users/signup", controller.UserSignUp())
	incomingRoutes.POST("users/login", controller.UserLogIn())

	// customer authentication
	incomingRoutes.POST("customers/signup", controller.CustomerSignUp())
	incomingRoutes.POST("customers/login", controller.CustomerLogIn())
}

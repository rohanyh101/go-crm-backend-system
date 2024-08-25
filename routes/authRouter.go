package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rohanhonnakatti/golang-jwt-auth/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.UserSignUp())
	incomingRoutes.POST("users/login", controller.UserLogIn())

	incomingRoutes.POST("customers/signup", controller.CustomerSignUp())
	incomingRoutes.POST("customers/login", controller.CustomerLogIn())
}

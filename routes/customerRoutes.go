package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rohanhonnakatti/golang-jwt-auth/controllers"
	"github.com/rohanhonnakatti/golang-jwt-auth/middleware"
)

func CustomerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.AuthenticateCustomer())

	incomingRoutes.GET("/customers", controller.GetCustomers())
	incomingRoutes.GET("/customers/:customer_id", controller.GetCustomer())
	incomingRoutes.PUT("/customers/:customer_id", controller.UpdateCustomer())
	incomingRoutes.DELETE("/customers/:customer_id", controller.DeleteCustomer())

	incomingRoutes.GET("/customers/ticket/:user_id", controller.GetTicketsByUserID())
	incomingRoutes.POST("/customers/ticket/", controller.CreateTicketAndSendEmail())
	incomingRoutes.PUT("/customers/ticket/:ticket_id", controller.UpdateTicket())
	incomingRoutes.DELETE("/customers/ticket/:ticket_id", controller.DeleteTicket())
}

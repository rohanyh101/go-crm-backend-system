package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/roh4nyh/matrice_ai/controllers"
	"github.com/roh4nyh/matrice_ai/middleware"
)

func CustomerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.AuthenticateCustomer())

	// customer crud operations
	incomingRoutes.GET("/customers", controller.GetCustomers())
	incomingRoutes.GET("/customers/:customer_id", controller.GetCustomer())
	incomingRoutes.PUT("/customers/:customer_id", controller.UpdateCustomer())
	incomingRoutes.DELETE("/customers/:customer_id", controller.DeleteCustomer())

	// customer services
	// get all tickets, only for admin
	incomingRoutes.GET("/customers/tickets/", controller.GetAllTickets())

	// create ticket
	incomingRoutes.POST("/customers/ticket/:interaction_id", controller.CreateTicket())

	// get all tickets related to current user
	incomingRoutes.GET("/customers/ticket/:user_id", controller.GetTicketsByUserID())

	// update ticket
	incomingRoutes.PUT("/customers/ticket/:ticket_id", controller.UpdateTicket())

	// delete ticket
	incomingRoutes.DELETE("/customers/ticket/:ticket_id", controller.DeleteTicket())
}

package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rohanhonnakatti/golang-jwt-auth/database"
	"github.com/rohanhonnakatti/golang-jwt-auth/models"
	"github.com/rohanhonnakatti/golang-jwt-auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	TicketDatabaseName   = "Cluster0"
	TicketCollectionName = "tickets"
)

var TicketValidate = validator.New()
var TicketCollection *mongo.Collection = database.OpenCollection(TicketDatabaseName, TicketCollectionName)

func CreateTicketAndSendEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var ticket models.Ticket
		if err := c.BindJSON(&ticket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := userValidate.Struct(ticket)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		ticket.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		ticket.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		ticket.ID = primitive.NewObjectID()
		ticket.TicketId = ticket.ID.Hex()

		resultInsertionNumber, insertErr := UserCollection.InsertOne(ctx, ticket)
		if insertErr != nil {
			msg := fmt.Sprintln("fialed to create Interaction")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		email := c.GetString("email")
		// !!! where to send...
		go utils.SendEmailToUser(ticket, email)

		c.JSON(http.StatusCreated, resultInsertionNumber)
	}
}

func UpdateTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketId := c.Param("ticket_id")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var ticket models.Ticket

		if err := c.BindJSON(&ticket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateObj := bson.M{}

		if ticket.Status != nil {
			updateObj["status"] = ticket.Status
		}

		if ticket.Description != nil {
			updateObj["description"] = ticket.Description
		}

		// if user.Role != nil && *user.Role == "ADMIN" {
		// 	updateObj["role"] = user.Role
		// }

		updateObj["updatedat"] = time.Now()

		filter := bson.M{"ticketid": bson.M{"$eq": ticketId}}
		update := bson.M{"$set": updateObj}

		result, err := UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating user"})
			return
		}

		c.JSON(http.StatusOK, result)

	}
}

func GetTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		// only admin can access all tickets !!!
		// if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		var tickets []models.Ticket

		cursor, err := UserCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing users"})
			return
		}

		if err = cursor.All(ctx, &tickets); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding user data"})
			return
		}

		if len(tickets) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": "no tickets available"})
			return
		}

		c.JSON(http.StatusOK, tickets)
	}
}

// admin feature !!!
// func GetTicketByID() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 		defer cancel()

// 		ticketId := c.Param("ticketId")

// 		var ticket models.Ticket
// 		err := TicketCollection.FindOne(ctx, bson.M{"ticketid": ticketId}).Decode(&ticket)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching ticket"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, ticket)
// 	}
// }

func GetTicketsByUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var tickets []models.Ticket
		cursor, err := InteractionCollection.Find(ctx, bson.M{"interactionid": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing users"})
			return
		}

		if err = cursor.All(ctx, &tickets); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding user data"})
			return
		}

		if len(tickets) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": "no tickets raised by this user"})
			return
		}

		c.JSON(http.StatusOK, tickets)

	}
}

func DeleteTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketId := c.Param("ticket_id")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		_, err := TicketCollection.DeleteOne(ctx, bson.M{"ticketid": ticketId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting ticket"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ticket deleted successfully"})
	}
}

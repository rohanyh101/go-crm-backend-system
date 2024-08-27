package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/roh4nyh/matrice_ai/database"
	"github.com/roh4nyh/matrice_ai/models"
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

func CreateTicket() gin.HandlerFunc {
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

		interactionIdStr := c.Param("interaction_id")
		interactionId, err := primitive.ObjectIDFromHex(interactionIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interaction ID"})
			return
		}

		customerIdStr := c.GetString("cid")
		customerId, err := primitive.ObjectIDFromHex(customerIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
			return
		}

		filter := bson.M{
			"_id":         interactionId,
			"customer_id": customerId,
		}

		var interaction models.Interaction
		err = InteractionCollection.FindOne(ctx, filter).Decode(&interaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "customer not belongs to this interaction or interaction not exists"})
			return
		}

		ticket.CustomerID = customerId
		ticket.InteractionID = interactionId
		ticket.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		ticket.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		ticket.ID = primitive.NewObjectID()
		ticket.TicketId = ticket.ID.Hex()

		resultInsertionNumber, insertErr := TicketCollection.InsertOne(ctx, ticket)
		if insertErr != nil {
			msg := fmt.Sprintln("fialed to create Interaction")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusCreated, resultInsertionNumber)
	}
}

func UpdateTicket() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var ticket models.Ticket
		if err := c.BindJSON(&ticket); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ticketIdStr := c.Param("ticket_id")
		ticketId, err := primitive.ObjectIDFromHex(ticketIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
			return
		}

		customerIdStr := c.GetString("cid")
		customerId, err := primitive.ObjectIDFromHex(customerIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
			return
		}

		filter := bson.M{
			"_id":         ticketId,
			"customer_id": customerId,
		}

		// var ticket models.Ticket
		err = TicketCollection.FindOne(ctx, filter).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while fetching interaction"})
			return
		}

		updateObj := bson.M{}

		if ticket.Status != nil {
			updateObj["status"] = ticket.Status
		}

		if ticket.Description != nil {
			updateObj["description"] = ticket.Description
		}

		updateObj["updated_at"] = time.Now()

		update := bson.M{"$set": updateObj}

		_, err = TicketCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating ticket"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ticket updated successfully"})

	}
}

func GetAllTickets() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var tickets []models.Ticket

		cursor, err := TicketCollection.Find(ctx, bson.M{})
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
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		userIdStr := c.Param("user_id")
		userId, err := primitive.ObjectIDFromHex(userIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
			return
		}

		var tickets []models.Ticket
		cursor, err := TicketCollection.Find(ctx, bson.M{"user_id": userId})
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
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		ticketIdStr := c.Param("ticket_id")
		ticketId, err := primitive.ObjectIDFromHex(ticketIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
			return
		}

		customerIdStr := c.GetString("cid")
		customerId, err := primitive.ObjectIDFromHex(customerIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
			return
		}

		filter := bson.M{
			"_id":         ticketId,
			"customer_id": customerId,
		}

		_, err = TicketCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ticket deletion failed or ticket not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ticket deleted successfully"})
	}
}

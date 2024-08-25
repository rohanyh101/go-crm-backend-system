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
	InteractionDatabaseName   = "Cluster0"
	InteractionCollectionName = "interactions"
)

var InteractionValidate = validator.New()
var InteractionCollection *mongo.Collection = database.OpenCollection(InteractionDatabaseName, InteractionCollectionName)

func CreateInteraction() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var interaction models.Interaction
		if err := c.BindJSON(&interaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := userValidate.Struct(interaction)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		interaction.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		interaction.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		interaction.ID = primitive.NewObjectID()
		interaction.InteractionId = interaction.ID.Hex()

		userIDStr := c.GetString("uid")
		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		interaction.UserID = userID

		resultInsertionNumber, insertErr := UserCollection.InsertOne(ctx, interaction)
		if insertErr != nil {
			msg := fmt.Sprintln("fialed to create Interaction")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		email := c.GetString("email")

		var user models.User
		err = UserCollection.FindOne(ctx, bson.M{"user_id": interaction.UserID}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		go utils.SendEmailInteraction(interaction, email)

		// go sendEmailToClient(interaction, *user.Email)

		c.JSON(http.StatusCreated, resultInsertionNumber)
	}
}

// this is admin feature !!!
// func GetInteractions() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
// 		defer cancel()

// 		// if err := helper.CheckUserType(c, "ADMIN"); err != nil {
// 		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		// 	return
// 		// }

// 		var interactions []models.Interaction

// 		cursor, err := InteractionCollection.Find(ctx, bson.M{})
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing users"})
// 			return
// 		}

// 		if err = cursor.All(ctx, &interactions); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding user data"})
// 			return
// 		}

// 		if len(interactions) == 0 {
// 			c.JSON(http.StatusOK, gin.H{"error": "no interactions available"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, interactions)
// 	}
// }

func GetInteractionsByUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("uid")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var interactions []models.Interaction

		cursor, err := InteractionCollection.Find(ctx, bson.M{"userid": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing users"})
			return
		}

		if err = cursor.All(ctx, &interactions); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding user data"})
			return
		}

		if len(interactions) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": "no users available"})
			return
		}

		c.JSON(http.StatusOK, interactions)
	}
}

func DeleteInteraction() gin.HandlerFunc {
	return func(c *gin.Context) {
		interactionId := c.Param("interaction_id")

		// need to do user himself can only delete his interaction, not others ???
		// if err := helper.MatchUserTypeToUid(c, interactionId); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		result, err := UserCollection.DeleteOne(ctx, bson.M{"interactionid": interactionId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting user"})
			return
		}

		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Interaction not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

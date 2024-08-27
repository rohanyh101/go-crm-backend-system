package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/roh4nyh/matrice_ai/database"
	helper "github.com/roh4nyh/matrice_ai/helpers"
	"github.com/roh4nyh/matrice_ai/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	customerdatabaseName   = "Cluster0"
	customerCollectionName = "customers"
)

var customerValidate = validator.New()
var CustomerCollection *mongo.Collection = database.OpenCollection(customerdatabaseName, customerCollectionName)

func CustomerSignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var customer models.Customer
		if err := c.BindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := customerValidate.Struct(customer)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := CustomerCollection.CountDocuments(ctx, bson.M{"email": customer.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking for email"})
			log.Panic(err)
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exists"})
			return
		}

		password := HashPassword(*customer.Password)
		customer.Password = &password

		customer.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		customer.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		customer.ID = primitive.NewObjectID()
		customer.CustomerId = customer.ID.Hex()

		token, _ := helper.GenerateCustomerToken(*customer.Email, *customer.Name, customer.CustomerId)
		customer.Token = &token

		resultInsertionNumber, insertErr := CustomerCollection.InsertOne(ctx, customer)
		if insertErr != nil {
			msg := fmt.Sprintln("Customer item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusCreated, resultInsertionNumber)
	}
}

func CustomerLogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var customer models.Customer
		var foundCustomer models.Customer

		if err := c.BindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := CustomerCollection.FindOne(ctx, bson.M{"email": customer.Email}).Decode(&foundCustomer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*customer.Password, *foundCustomer.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if foundCustomer.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "customer not found"})
			return
		}

		token, err := helper.GenerateCustomerToken(*foundCustomer.Email, *foundCustomer.Name, foundCustomer.CustomerId)
		if err != nil || token == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		helper.UpdateCustomerToken(token, foundCustomer.CustomerId)

		err = CustomerCollection.FindOne(ctx, bson.M{"customer_id": foundCustomer.CustomerId}).Decode(&foundCustomer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, foundCustomer)
	}
}

func GetCustomers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var customers []models.Customer

		cursor, err := CustomerCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while listing customers"})
			return
		}

		if err = cursor.All(ctx, &customers); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while decoding customer data"})
			return
		}

		if len(customers) == 0 {
			c.JSON(http.StatusOK, gin.H{"error": "no customers available"})
			return
		}

		c.JSON(http.StatusOK, customers)
	}
}

func GetCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId := c.Param("customer_id")

		if err := helper.MatchCustomerTypeToCid(c, customerId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var customer models.Customer
		err := CustomerCollection.FindOne(ctx, bson.M{"customer_id": customerId}).Decode(&customer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if customer.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "customer not found"})
			return
		}

		c.JSON(http.StatusOK, customer)
	}
}

func UpdateCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId := c.Param("customer_id")

		if err := helper.MatchCustomerTypeToCid(c, customerId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var customer models.Customer

		if err := c.BindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateObj := bson.M{}

		if customer.Name != nil {
			updateObj["name"] = customer.Name
		}

		if customer.Email != nil {
			updateObj["email"] = customer.Email
		}

		if customer.Company != nil {
			updateObj["company"] = customer.Company
		}

		if customer.Phone != nil {
			updateObj["phone"] = customer.Phone
		}

		if customer.Password != nil {
			password := HashPassword(*customer.Password)
			updateObj["password"] = password
		}

		updateObj["updated_at"] = time.Now()

		filter := bson.M{"customer_id": bson.M{"$eq": customerId}}
		update := bson.M{"$set": updateObj}

		_, err := CustomerCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while updating customer"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "customer updated successfully"})
	}
}

func DeleteCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId := c.Param("customer_id")

		if err := helper.MatchCustomerTypeToCid(c, customerId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.M{"customer_id": bson.M{"$eq": customerId}}

		_, err := CustomerCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting customer"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
	}
}

package helpers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rohanhonnakatti/golang-jwt-auth/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedCustomerDetails struct {
	Email string
	Name  string
	Cid   string
	jwt.StandardClaims
}

const (
	customerDatabaseName   = "Cluster0"
	customerCollectionName = "customers"
)

var CustomerCollection *mongo.Collection = database.OpenCollection(customerDatabaseName, customerCollectionName)
var CUSTOMER_SECRET_KEY string = os.Getenv("CUSTOMER_SECRET_KEY")

func GenerateCustomerToken(email, name, cId string) (signedToken string, err error) {
	claims := &SignedCustomerDetails{
		Email: email,
		Name:  name,
		Cid:   cId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(CUSTOMER_SECRET_KEY))
	if err != nil {
		msg := fmt.Sprintf("Error signing Token: %v", err)
		return "", errors.New(msg)
	}

	return token, nil
}

func UpdateCustomerToken(signedToken, customerId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updatedat", Value: updatedAt})

	upsert := true
	filter := bson.M{"customerid": customerId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := UserCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)

	if err != nil {
		log.Panic(err)
		return
	}

	// return
}

func ValidateCustomerToken(signedToken string) (claims *SignedCustomerDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedCustomerDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(CUSTOMER_SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedCustomerDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("the token has expired")
		msg = err.Error()
		return
	}
	return claims, msg
}

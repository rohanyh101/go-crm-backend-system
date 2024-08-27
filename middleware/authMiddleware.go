package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/roh4nyh/matrice_ai/helpers"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header found"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateUserToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("name", claims.Name)
		c.Set("role", claims.Role)
		c.Set("uid", claims.Uid)

		c.Next()
	}
}

func AuthenticateCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header found"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateCustomerToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("cid", claims.Cid)
		c.Set("email", claims.Email)
		c.Set("name", claims.Name)

		c.Next()
	}
}

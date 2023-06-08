package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authenticate implements Bearer token authentication
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorizationHeader := c.Request.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.JSON(400, gin.H{
				"status":  0,
				"message": "invalid token",
			})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}
			jwtSecret := os.Getenv("JWT_SECRET")
			return jwtSecret, nil
		})

		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"status":  0,
				"message": "invalid token",
			})
			return
		}
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(400, gin.H{
				"status":  0,
				"message": "invalid token",
			})
			return
		}
		c.Next()
	}
}

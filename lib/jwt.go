package lib

import (
	"crypto/sha512"
	"crypto/subtle"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/far00kaja/learn-go-with-case/internal/auth/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JwtClaims struct {
	Id uuid.UUID `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(authId uuid.UUID) (string, error) {

	JWT_SECRET := os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Fatalln("missing JWT_SECRET from env")
	}
	signKey := []byte(JWT_SECRET)

	// Create the Claims
	claims := JwtClaims{
		authId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signKey)

	if err != nil {
		fmt.Println(err.Error())
		return signedToken, err
	}
	return signedToken, nil
}

func VerifyJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerPrefix := "Bearer"
		headerToken := strings.Split(ctx.GetHeader("Authorization"), " ")

		if !SecureCompare(bearerPrefix, headerToken[0]) {
			ctx.AbortWithStatusJSON(401,
				gin.H{"code": 0,
					"message": "unauthorized"})
		}
		// idToken := strings.TrimSpace(strings.Replace(headerToken, "Bearer", "", 1))
		fmt.Println("idToken", headerToken[1])
		token, err := jwt.ParseWithClaims(headerToken[1], &dto.TokensResponse{}, func(token *jwt.Token) (interface{}, error) {
			jwtSecret := os.Getenv("JWT_SECRET")
			return []byte(jwtSecret), nil
		})
		if claims, ok := token.Claims.(*dto.TokensResponse); ok && token.Valid {
			fmt.Println(claims.Token)
			ctx.Next()
			return
		}
		if err != nil {
			fmt.Println("gagal")
			ctx.AbortWithStatusJSON(401, gin.H{
				"status":  0,
				"message": "unauthorized",
			})
			return
		}
	}
}

func MiddlewareCheckRedis() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(headerToken, "Bearer", "", 1))
		fmt.Println(idToken)
		ctx.Next()
		// return
		// result,err:= lib.GetRedisFromKeyStrValue(c.)
	}
}

func SecureCompare(given string, actual string) bool {
	givenSha := sha512.Sum512([]byte(given))
	actualSha := sha512.Sum512([]byte(actual))

	return subtle.ConstantTimeCompare(givenSha[:], actualSha[:]) == 1
}

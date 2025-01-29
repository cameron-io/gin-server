package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	gin_jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	IdentityKey = "identity"
)

func InitHandlerMiddleware(authMiddleware *gin_jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := gin_jwt.ExtractClaims(c)
	return claims[IdentityKey]
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func CreateAuthToken(email string) (string, error) {
	method := &jwt.SigningMethodHMAC{
		Name: jwt.SigningMethodHS256.Name,
		Hash: jwt.SigningMethodHS256.Hash,
	}
	token := jwt.NewWithClaims(
		method,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserIdFromClaims(c *gin.Context) string {
	user, _ := c.Get(IdentityKey)
	return user.(map[string]interface{})["id"].(string)
}

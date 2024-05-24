package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TokenAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token_str, cookie_err := ctx.Cookie("token")
		if cookie_err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": cookie_err.Error()})
			return
		}
		key_fun := func(token *jwt.Token) (interface{}, error) {
			// validate JWT-signing algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		}
		token, err := jwt.Parse(token_str, key_fun)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("user_claims", claims)
		} else {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"msg": "invalid token."})
		}

		ctx.Next()
	}
}

package middleware

import (
	"log"

	"cameron.io/gin-server/api/dto"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	IdentityKey = "identity"
)

func InitHandlerMiddleware(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func GetUserIdFromClaims(c *gin.Context) string {
	user, _ := c.Get(IdentityKey)
	return user.(map[string]interface{})["id"].(string)
}

func PayloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if user, ok := data.(*dto.Identity); ok {
			return jwt.MapClaims{
				IdentityKey: user,
			}
		}
		return jwt.MapClaims{}
	}
}

func IdentityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return claims[IdentityKey]
	}
}

package middleware

import (
	"log"
	"os"
	"time"

	"cameron.io/gin-server/models"
	"cameron.io/gin-server/services"
	"cameron.io/gin-server/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var (
	identityKey = "identity"
)

func InitParams() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       os.Getenv("SERVER_NAME") + "_user",
		Key:         []byte(os.Getenv("JWT_SECRET")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc(),

		IdentityHandler: identityHandler(),
		Authenticator:   authHandler(),

		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	}
}

func InitHandlerMiddleware(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if user, ok := data.(*models.Identity); ok {
			return jwt.MapClaims{
				identityKey: user,
			}
		}
		return jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return claims[identityKey]
	}
}

func authHandler() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var user_auth models.Auth

		if err := c.ShouldBindJSON(&user_auth); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		if err := validator.New().Struct(user_auth); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		existing_user, db_err := services.FindUserByEmail(c, user_auth.Email)
		if db_err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		if err := utils.MatchPasswords(
			user_auth.Password,
			existing_user["password"].(string),
		); err != nil {
			return "", err
		}

		return &models.Identity{
			Name:   existing_user["name"].(string),
			Email:  existing_user["email"].(string),
			Avatar: existing_user["avatar"].(string),
		}, nil
	}
}

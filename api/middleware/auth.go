package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"cameron.io/gin-server/api/dto"
	"cameron.io/gin-server/services"
	"cameron.io/gin-server/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

		SendCookie:     true,
		SecureCookie:   os.Getenv("SERVER_ENV") == "production",
		CookieHTTPOnly: true,
		CookieDomain:   os.Getenv("SERVER_URI"),
		CookieName:     "token",
		TokenLookup:    "cookie:token",
		CookieSameSite: http.SameSiteDefaultMode, //SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode, SameSiteNoneMode

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
		if user, ok := data.(*dto.Identity); ok {
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
		var userLogin dto.Login

		if err := c.ShouldBindJSON(&userLogin); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		if err := validator.New().Struct(userLogin); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		existingUser, dbErr := services.FindUserByEmail(c, userLogin.Email)
		if dbErr != nil {
			return "", jwt.ErrFailedAuthentication
		}

		if err := utils.MatchPasswords(
			userLogin.Password,
			existingUser["password"].(string),
		); err != nil {
			return "", err
		}

		return &dto.Identity{
			Id:     existingUser["_id"].(primitive.ObjectID).Hex(),
			Name:   existingUser["name"].(string),
			Email:  existingUser["email"].(string),
			Avatar: existingUser["avatar"].(string),
		}, nil
	}
}

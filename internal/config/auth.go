package config

import (
	"net/http"
	"os"
	"time"

	"cameron.io/gin-server/internal/dto"
	"cameron.io/gin-server/internal/handlers"
	"cameron.io/gin-server/pkg/auth"
	gin_jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitParams(handler handlers.AuthHandler) *gin_jwt.GinJWTMiddleware {
	return &gin_jwt.GinJWTMiddleware{
		Realm:       os.Getenv("SERVER_NAME") + "_user",
		Key:         []byte(os.Getenv("JWT_SECRET")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: auth.IdentityKey,
		PayloadFunc: payloadFunc(),
		KeyFunc:     auth.KeyFunc,

		IdentityHandler: identityHandler(),
		Authenticator:   handler.Authenticator,

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

func payloadFunc() func(data interface{}) gin_jwt.MapClaims {
	return func(data interface{}) gin_jwt.MapClaims {
		if user, ok := data.(*dto.Identity); ok {
			return gin_jwt.MapClaims{
				auth.IdentityKey: user,
			}
		}
		return gin_jwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := gin_jwt.ExtractClaims(c)
		return claims[auth.IdentityKey]
	}
}

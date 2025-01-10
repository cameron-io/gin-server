package config

import (
	"net/http"
	"os"
	"time"

	"cameron.io/gin-server/api/controllers"
	"cameron.io/gin-server/api/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func InitParams(controller controllers.AuthController) *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       os.Getenv("SERVER_NAME") + "_user",
		Key:         []byte(os.Getenv("JWT_SECRET")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: middleware.IdentityKey,
		PayloadFunc: middleware.PayloadFunc(),

		IdentityHandler: middleware.IdentityHandler(),
		Authenticator:   controller.Authenticator,

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

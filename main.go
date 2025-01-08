package main

import (
	"log"
	"os"

	"cameron.io/gin-server/api/controllers"
	"cameron.io/gin-server/api/middleware"
	"cameron.io/gin-server/application/services"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	rGroupApi := r.Group("/api")

	// Accounts - services
	userService := services.NewUserService()
	authService := services.NewAuthService(userService)

	// Accounts - middleware
	authHandle, err := jwt.New(middleware.InitParams(authService))
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	r.Use(middleware.InitHandlerMiddleware(authHandle))

	// Accounts - controller
	controllers.NewUserController(rGroupApi, authHandle, userService)

	// Profiles
	profileService := services.NewProfileService()
	controllers.NewProfileController(rGroupApi, authHandle, profileService)

	r.SetTrustedProxies(nil)
	r.Run(os.Getenv("SERVER_URI"))
}

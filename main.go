package main

import (
	"log"
	"os"
	"time"

	"cameron.io/gin-server/api/config"
	"cameron.io/gin-server/api/controllers"
	"cameron.io/gin-server/api/middleware"
	"cameron.io/gin-server/domain/services"
	"cameron.io/gin-server/infra/db/mongo/repositories"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://github.com"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	rGroupApi := r.Group("/api")

	// Accounts - services
	userRepository := repositories.NewGenRepository("user")
	profileRepository := repositories.NewGenRepository("profile")

	userService := services.NewUserService(userRepository, profileRepository)

	// Accounts - middleware
	authService := services.NewAuthService(userService)
	authController := controllers.NewAuthController(authService)
	authHandle, err := jwt.New(config.InitParams(*authController))
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	r.Use(middleware.InitHandlerMiddleware(authHandle))

	// Accounts - controller
	controllers.NewUserController(rGroupApi, authHandle, userService)

	// Profiles
	profileService := services.NewProfileService(profileRepository)
	controllers.NewProfileController(rGroupApi, authHandle, profileService)

	r.SetTrustedProxies(nil)
	r.Run(os.Getenv("SERVER_URI"))
}

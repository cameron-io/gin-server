package main

import (
	"log"
	"os"
	"time"

	"cameron.io/gin-server/internal/api/config"
	"cameron.io/gin-server/internal/api/handlers"
	"cameron.io/gin-server/internal/api/middleware"
	"cameron.io/gin-server/internal/domain/services"
	"cameron.io/gin-server/pkg/db/mongo/repositories"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://studio.apollographql.com"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
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
	authHandler := handlers.NewAuthHandler(userService)
	authHandle, err := jwt.New(config.InitParams(*authHandler))
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	r.Use(middleware.InitHandlerMiddleware(authHandle))

	// Accounts - request handler
	mailService := services.NewMailService()
	handlers.NewUserHandler(rGroupApi, authHandle, userService, mailService)

	// Profiles - request handler
	profileService := services.NewProfileService(profileRepository)
	handlers.NewProfileHandler(rGroupApi, authHandle, profileService)

	// Products - GraphQL handler
	handlers.NewGQueryHandler(rGroupApi)

	r.SetTrustedProxies(nil)
	r.Run(os.Getenv("SERVER_URI"))
}

package main

import (
	"log"
	"os"

	"cameron.io/gin-server/api"
	"cameron.io/gin-server/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// register middleware
	authHandle, err := jwt.New(middleware.InitParams())
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	r.Use(middleware.InitHandlerMiddleware(authHandle))

	rGroupApi := r.Group("/api")

	// Accounts - Public Routes
	rGroupAcc := rGroupApi.Group("/accounts")
	rGroupAcc.POST("/register", api.RegisterUser)
	rGroupAcc.POST("/login", authHandle.LoginHandler)

	// Accounts - Protected Routes
	rGroupAuth := rGroupAcc.Group("/", authHandle.MiddlewareFunc())
	rGroupAuth.GET("/refresh_token", authHandle.RefreshHandler)
	rGroupAuth.POST("/logout", authHandle.LogoutHandler)
	rGroupAuth.GET("/info", api.GetUserInfo)
	rGroupAuth.DELETE("/", api.DeleteUser)

	// Profiles - Protected Routes
	rGroupProfile := rGroupApi.Group("/profiles", authHandle.MiddlewareFunc())
	rGroupProfile.POST("/", api.UpsertProfile)
	rGroupProfile.GET("/", api.GetAllProfiles)
	rGroupProfile.GET("/me", api.GetCurrentUserProfile)
	rGroupProfile.GET("/user/:user_id", api.GetProfileByUserId)

	r.SetTrustedProxies(nil)
	r.Run(os.Getenv("SERVER_URI"))
}

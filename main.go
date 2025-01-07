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

	rGroupAcc := r.Group("/api/accounts")
	rGroupAcc.POST("/register", api.RegisterUser)
	rGroupAcc.POST("/login", authHandle.LoginHandler)

	// Protected Routes
	rGroupAuth := rGroupAcc.Group("/", authHandle.MiddlewareFunc())
	rGroupAuth.GET("/refresh_token", authHandle.RefreshHandler)
	rGroupAuth.POST("/logout", authHandle.LogoutHandler)
	rGroupAuth.GET("/test", testHandler)

	r.SetTrustedProxies(nil)
	r.Run(os.Getenv("SERVER_URI"))
}

func testHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// user, _ := c.Get("identity")
	c.JSON(200, gin.H{
		"claims": claims,
	})
}

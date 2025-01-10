package controllers

import (
	"net/http"
	"time"

	"cameron.io/gin-server/api/middleware"
	"cameron.io/gin-server/domain/entities"
	"cameron.io/gin-server/domain/interfaces"
	"cameron.io/gin-server/domain/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(
	r *gin.RouterGroup,
	authHandle *jwt.GinJWTMiddleware,
	service interfaces.UserService,
) {
	controller := &UserController{
		service: service,
	}
	// Accounts - Public Routes
	rGroupAcc := r.Group("/accounts")
	rGroupAcc.POST("/register", controller.RegisterUser)
	rGroupAcc.POST("/login", authHandle.LoginHandler)

	// Accounts - Protected Routes
	rGroupAuth := rGroupAcc.Group("/", authHandle.MiddlewareFunc())
	rGroupAuth.GET("/refresh_token", authHandle.RefreshHandler)
	rGroupAuth.POST("/logout", authHandle.LogoutHandler)
	rGroupAuth.GET("/info", controller.GetUserInfo)
	rGroupAuth.DELETE("/", controller.DeleteUser, authHandle.LogoutHandler)
}

func (uc *UserController) GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("identity")
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var newUser entities.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate the User struct
	if err := validator.New().Struct(newUser); err != nil {
		// Validation failed, handle the error
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	existingUser, err := uc.service.FindUserByEmail(ctx, newUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
		return
	}
	if existingUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "user already exists"})
		return
	}

	// Hash password
	if newUser.Password, err = utils.HashPassword(newUser.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"unexpected_error": err.Error()})
		return
	}

	// Create new user
	newUser.CreatedAt = time.Now().UnixMilli()

	if err := uc.service.CreateUser(ctx, newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"db_create_error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	userId := middleware.GetUserIdFromClaims(ctx)

	res, err := uc.service.DeleteUserByID(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !res {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}

package handlers

import (
	"net/http"
	"time"

	"cameron.io/gin-server/api/dto"
	"cameron.io/gin-server/api/middleware"
	"cameron.io/gin-server/domain/entities"
	"cameron.io/gin-server/domain/interfaces"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserHandler struct {
	userService interfaces.UserService
	mailService interfaces.MailService
}

func NewUserHandler(
	r *gin.RouterGroup,
	authHandle *jwt.GinJWTMiddleware,
	userService interfaces.UserService,
	mailService interfaces.MailService,
) {
	userHandle := &UserHandler{
		userService: userService,
		mailService: mailService,
	}
	// Accounts - Public Routes
	rGroupAcc := r.Group("/accounts")
	rGroupAcc.POST("/register", userHandle.RegisterUser)
	rGroupAcc.POST("/login", userHandle.LoginUser)
	rGroupAcc.GET("/login", authHandle.LoginHandler)

	// Accounts - Protected Routes
	rGroupAuth := rGroupAcc.Group("/", authHandle.MiddlewareFunc())
	rGroupAuth.GET("/refresh_token", authHandle.RefreshHandler)
	rGroupAuth.POST("/logout", authHandle.LogoutHandler)
	rGroupAuth.GET("/info", userHandle.GetUserInfo)
	rGroupAuth.DELETE("/", userHandle.DeleteUser, authHandle.LogoutHandler)
}

func (u *UserHandler) GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("identity")
	ctx.JSON(http.StatusOK, user)
}

func (u *UserHandler) LoginUser(ctx *gin.Context) {
	var userLogin dto.Login

	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate the User struct
	if err := validator.New().Struct(userLogin); err != nil {
		// Validation failed, handle the error
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, tokenErr := middleware.CreateAuthToken(userLogin.Email)
	if tokenErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": tokenErr})
		return
	}

	u.mailService.MailLoginToken(token)

	ctx.JSON(http.StatusOK, gin.H{"msg": "Check your inbox!"})
}

func (u *UserHandler) RegisterUser(ctx *gin.Context) {
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
	existingUser, err := u.userService.FindUserByEmail(ctx, newUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
		return
	}
	if existingUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "user already exists"})
		return
	}

	// Create new user
	newUser.CreatedAt = time.Now().UnixMilli()

	if err := u.userService.CreateUser(ctx, newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"db_create_error": err.Error()})
		return
	}

	token, tokenErr := middleware.CreateAuthToken(newUser.Email)
	if tokenErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": tokenErr})
		return
	}

	u.mailService.MailLoginToken(token)

	ctx.Status(http.StatusCreated)
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	userId := middleware.GetUserIdFromClaims(ctx)

	res, err := u.userService.DeleteUserByID(ctx, userId)
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

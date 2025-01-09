package controllers

import (
	"net/http"

	"cameron.io/gin-server/api/middleware"
	"cameron.io/gin-server/application/interfaces"
	"cameron.io/gin-server/domain/entities"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ProfileController struct {
	service interfaces.ProfileService
}

func NewProfileController(
	rGroupApi *gin.RouterGroup,
	authHandle *jwt.GinJWTMiddleware,
	service interfaces.ProfileService,
) {
	controller := &ProfileController{
		service: service,
	}
	rGroupProfile := rGroupApi.Group("/profiles", authHandle.MiddlewareFunc())
	rGroupProfile.POST("/", controller.UpsertProfile)
	rGroupProfile.GET("/", controller.GetAllProfiles)
	rGroupProfile.GET("/me", controller.GetCurrentUserProfile)
	rGroupProfile.GET("/user/:user_id", controller.GetProfileByUserId)
}

func (pc *ProfileController) GetCurrentUserProfile(ctx *gin.Context) {
	userId := middleware.GetUserIdFromClaims(ctx)

	profile, dbErr := pc.service.GetProfileByUserId(ctx, userId)
	if dbErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

func (pc *ProfileController) GetProfileByUserId(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	profile, dbErr := pc.service.GetProfileByUserId(ctx, userId)
	if dbErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

func (pc *ProfileController) UpsertProfile(ctx *gin.Context) {
	var newProfile entities.Profile

	if err := ctx.ShouldBindJSON(&newProfile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.New().Struct(newProfile); err != nil {
		// Validation failed, handle the error
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := middleware.GetUserIdFromClaims(ctx)

	profile, err := pc.service.UpsertProfile(ctx, userId, newProfile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"db_upsert_error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, profile)
}

func (pc *ProfileController) GetAllProfiles(ctx *gin.Context) {
	profiles, err := pc.service.GetAllProfiles(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"db_all_profiles_error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, profiles)
}

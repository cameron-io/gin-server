package handlers

import (
	"net/http"

	"cameron.io/gin-server/internal/domain/include"
	"cameron.io/gin-server/internal/domain/models"
	"cameron.io/gin-server/pkg/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ProfileHandler struct {
	service include.ProfileService
}

func NewProfileHandler(
	rGroupApi *gin.RouterGroup,
	authHandle *jwt.GinJWTMiddleware,
	service include.ProfileService,
) {
	handler := &ProfileHandler{
		service: service,
	}
	rGroupProfile := rGroupApi.Group("/profiles", authHandle.MiddlewareFunc())
	rGroupProfile.POST("/", handler.UpsertProfile)
	rGroupProfile.GET("/", handler.GetAllProfiles)
	rGroupProfile.GET("/me", handler.GetCurrentUserProfile)
	rGroupProfile.GET("/user/:user_id", handler.GetProfileByUserId)
}

func (pc *ProfileHandler) GetCurrentUserProfile(ctx *gin.Context) {
	userId := middleware.GetUserIdFromClaims(ctx)

	profile, dbErr := pc.service.GetProfileByUserId(ctx, userId)
	if dbErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

func (pc *ProfileHandler) GetProfileByUserId(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	profile, dbErr := pc.service.GetProfileByUserId(ctx, userId)
	if dbErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

func (pc *ProfileHandler) UpsertProfile(ctx *gin.Context) {
	var newProfile models.Profile

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

func (pc *ProfileHandler) GetAllProfiles(ctx *gin.Context) {
	profiles, err := pc.service.GetAllProfiles(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"db_all_profiles_error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, profiles)
}

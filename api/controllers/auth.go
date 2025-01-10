package controllers

import (
	"cameron.io/gin-server/api/dto"
	"cameron.io/gin-server/domain/interfaces"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AuthController struct {
	service interfaces.AuthService
}

func NewAuthController(
	service interfaces.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (uc *AuthController) Authenticator(ctx *gin.Context) (interface{}, error) {
	var userLogin dto.Login

	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	if err := validator.New().Struct(userLogin); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	return uc.service.Authenticator(ctx, userLogin.Email, userLogin.Password)
}

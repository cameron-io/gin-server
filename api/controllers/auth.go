package controllers

import (
	"encoding/json"

	"cameron.io/gin-server/api/dto"
	"cameron.io/gin-server/domain/entities"
	"cameron.io/gin-server/domain/interfaces"
	"cameron.io/gin-server/domain/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct {
	userService interfaces.UserService
}

func NewAuthController(
	userService interfaces.UserService) *AuthController {
	return &AuthController{
		userService: userService,
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

	existingUserObj, dbErr := uc.userService.FindUserByEmail(ctx, userLogin.Email)
	if dbErr != nil || existingUserObj == nil {
		return nil, jwt.ErrFailedAuthentication
	}

	var existingUser entities.User
	dbByte, _ := json.Marshal(existingUserObj)
	_ = json.Unmarshal(dbByte, &existingUser)

	if err := utils.MatchPasswords(
		userLogin.Password,
		existingUser.Password,
	); err != nil {
		return nil, err
	}

	// TODO: Extract based on DB_ENGINE var
	// MongoDB requires Id to be extracted via "_id" key
	dbUuidStr := existingUserObj["_id"].(primitive.ObjectID).Hex()

	return &dto.Identity{
		Id:     dbUuidStr,
		Name:   existingUser.Name,
		Email:  existingUser.Email,
		Avatar: existingUser.Avatar,
	}, nil
}

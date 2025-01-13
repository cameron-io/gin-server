package controllers

import (
	"encoding/json"

	"cameron.io/gin-server/api/dto"
	"cameron.io/gin-server/api/middleware"
	"cameron.io/gin-server/domain/entities"
	"cameron.io/gin-server/domain/interfaces"
	gin_jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func (uc *AuthController) Authenticator(ctx *gin.Context) (any, error) {
	token := ctx.Query("token")
	jwtToken, jwtErr := jwt.NewParser().Parse(token, middleware.KeyFunc)
	if jwtErr != nil {
		return nil, jwtErr
	}
	claims := gin_jwt.ExtractClaimsFromToken(jwtToken)

	userEmail := claims["email"].(string)
	existingUserObj, dbErr := uc.userService.FindUserByEmail(ctx, userEmail)
	if dbErr != nil || existingUserObj == nil {
		return nil, gin_jwt.ErrFailedAuthentication
	}

	var existingUser entities.User

	dbByte, byteErr := json.Marshal(existingUserObj)
	if byteErr != nil {
		return nil, byteErr
	}
	if err := json.Unmarshal(dbByte, &existingUser); err != nil {
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

package services

import (
	"encoding/json"

	"cameron.io/gin-server/api/dto"
	"cameron.io/gin-server/domain/entities"
	"cameron.io/gin-server/domain/interfaces"
	"cameron.io/gin-server/domain/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	userService interfaces.UserService
}

func NewAuthService(userService interfaces.UserService) interfaces.AuthService {
	return &AuthService{
		userService: userService,
	}
}

func (s *AuthService) Authenticator(ctx *gin.Context, userEmail string, userPassword string) (*dto.Identity, error) {
	existingUserObj, dbErr := s.userService.FindUserByEmail(ctx, userEmail)
	if dbErr != nil || existingUserObj == nil {
		return nil, jwt.ErrFailedAuthentication
	}

	var existingUser entities.User
	dbByte, _ := json.Marshal(existingUserObj)
	_ = json.Unmarshal(dbByte, &existingUser)

	if err := utils.MatchPasswords(
		userPassword,
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

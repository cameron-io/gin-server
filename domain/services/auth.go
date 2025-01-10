package services

import (
	"cameron.io/gin-server/api/dto"
	"cameron.io/gin-server/domain/interfaces"
	"cameron.io/gin-server/domain/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	service interfaces.UserService
}

func NewAuthService(service interfaces.UserService) interfaces.AuthService {
	return &AuthService{
		service: service,
	}
}

func (uc *AuthService) Authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var userLogin dto.Login

		if err := c.ShouldBindJSON(&userLogin); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		if err := validator.New().Struct(userLogin); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		existingUser, dbErr := uc.service.FindUserByEmail(c, userLogin.Email)
		if dbErr != nil || existingUser == nil {
			return "", jwt.ErrFailedAuthentication
		}

		if err := utils.MatchPasswords(
			userLogin.Password,
			existingUser["password"].(string),
		); err != nil {
			return "", err
		}

		return &dto.Identity{
			Id:     existingUser["_id"].(primitive.ObjectID).Hex(),
			Name:   existingUser["name"].(string),
			Email:  existingUser["email"].(string),
			Avatar: existingUser["avatar"].(string),
		}, nil
	}
}

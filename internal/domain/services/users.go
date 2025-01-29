package services

import (
	"cameron.io/gin-server/internal/domain/models"
	"cameron.io/gin-server/pkg/db/data"
	"cameron.io/gin-server/pkg/db/interfaces"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository    interfaces.GenRepository
	profileRepository interfaces.GenRepository
}

func NewUserService(
	userRepo interfaces.GenRepository,
	profileRepo interfaces.GenRepository) *UserService {
	return &UserService{
		userRepository:    userRepo,
		profileRepository: profileRepo,
	}
}

func (s *UserService) FindUserByEmail(c *gin.Context, email string) (data.Obj, error) {
	filter := map[string]any{
		"email": email,
	}
	result, err := s.userRepository.Find(c, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) CreateUser(c *gin.Context, new_user models.User) error {
	return s.userRepository.Insert(c, new_user)
}

func (s *UserService) DeleteUserByID(c *gin.Context, userId string) (bool, error) {
	uuid := data.StrToUuid(userId)
	profileFilter := map[string]any{"user": uuid}
	if _, err := s.profileRepository.Delete(c, profileFilter); err != nil {
		return false, err
	}
	userFilter := map[string]any{"_id": uuid}
	if _, err := s.userRepository.Delete(c, userFilter); err != nil {
		return false, err
	}
	return true, nil
}

package services

import (
	"cameron.io/gin-server/domain/entities"
	"cameron.io/gin-server/domain/interfaces"
	"cameron.io/gin-server/infra/data"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository    interfaces.GenRepository
	profileRepository interfaces.GenRepository
}

func NewUserService(
	userRepo interfaces.GenRepository,
	profileRepo interfaces.GenRepository) interfaces.UserService {
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

func (s *UserService) CreateUser(c *gin.Context, new_user entities.User) error {
	return s.userRepository.Insert(c, new_user)
}

func (s *UserService) DeleteUserByID(c *gin.Context, userId string) (bool, error) {
	id, conv_err := data.ConvToUuid(userId)
	if conv_err != nil {
		return false, conv_err
	}
	profileFilter := map[string]any{"user": id}
	if _, err := s.profileRepository.Delete(c, profileFilter); err != nil {
		return false, err
	}
	userFilter := map[string]any{"_id": id}
	if _, err := s.userRepository.Delete(c, userFilter); err != nil {
		return false, err
	}
	return true, nil
}

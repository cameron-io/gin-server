package services

import (
	"cameron.io/gin-server/internal/domain/include"
	"cameron.io/gin-server/internal/domain/models"
	"cameron.io/gin-server/pkg/db/data"
	"cameron.io/gin-server/pkg/db/interfaces"
	"github.com/gin-gonic/gin"
)

type ProfileService struct {
	repository interfaces.GenRepository
}

func NewProfileService(profileRepo interfaces.GenRepository) include.ProfileService {
	return &ProfileService{
		repository: profileRepo,
	}
}

func (s *ProfileService) GetProfileByUserId(
	c *gin.Context,
	userId string,
) (data.Obj, error) {
	uuid := data.StrToUuid(userId)
	filter := map[string]any{
		"user": uuid,
	}
	result, err := s.repository.Find(c, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProfileService) GetAllProfiles(c *gin.Context) ([]data.Obj, error) {
	results, err := s.repository.FindAll(c, 10)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *ProfileService) UpsertProfile(
	c *gin.Context,
	userId string,
	profile models.Profile,
) (data.Obj, error) {
	uuid := data.StrToUuid(userId)
	profile.User = uuid
	filter := map[string]any{"user": userId}
	return s.repository.Upsert(c, filter, profile)
}

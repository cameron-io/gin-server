package services

import (
	"cameron.io/gin-server/application/i_services"
	"cameron.io/gin-server/domain/data"
	"cameron.io/gin-server/domain/entities"
	"cameron.io/gin-server/domain/i_repositories"
	"cameron.io/gin-server/infra/db/mongo/repositories"
	"github.com/gin-gonic/gin"
)

type ProfileService struct {
	repository i_repositories.GenRepository
}

func NewProfileService() i_services.ProfileService {
	return &ProfileService{
		repository: repositories.NewGenRepository("profile"),
	}
}

func (s *ProfileService) GetProfileByUserId(
	c *gin.Context,
	userId string,
) (data.Obj, error) {
	filter := map[string]any{
		"user": userId,
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
	profile entities.Profile,
) (data.Obj, error) {
	id, err := data.ConvToUuid(userId)
	if err != nil {
		return nil, err
	}
	profile.User = id
	filter := map[string]any{"user": userId}
	return s.repository.Upsert(c, filter, profile)
}

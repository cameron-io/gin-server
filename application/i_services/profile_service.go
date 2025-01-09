package i_services

import (
	"cameron.io/gin-server/domain/data"
	"cameron.io/gin-server/domain/entities"
	"github.com/gin-gonic/gin"
)

type ProfileService interface {
	GetProfileByUserId(c *gin.Context, userId string) (data.Obj, error)
	GetAllProfiles(c *gin.Context) ([]data.Obj, error)
	UpsertProfile(c *gin.Context, userId string, profile entities.Profile) (data.Obj, error)
}

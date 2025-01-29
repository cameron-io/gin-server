package include

import (
	"cameron.io/gin-server/internal/domain/models"
	"cameron.io/gin-server/pkg/db/data"
	"github.com/gin-gonic/gin"
)

type ProfileService interface {
	GetProfileByUserId(c *gin.Context, userId string) (data.Obj, error)
	GetAllProfiles(c *gin.Context) ([]data.Obj, error)
	UpsertProfile(c *gin.Context, userId string, profile models.Profile) (data.Obj, error)
}

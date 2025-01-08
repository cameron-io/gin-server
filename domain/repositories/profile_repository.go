package repositories

import (
	"cameron.io/gin-server/domain/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ProfileRepository interface {
	GetProfileByUserId(c *gin.Context, userId string) (bson.M, error)
	GetAllProfiles(c *gin.Context) ([]bson.M, error)
	UpsertProfile(c *gin.Context, userId string, profile entities.Profile) (bson.M, error)
}

package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenericRepository interface {
	Upsert(c *gin.Context, entity interface{})
	FindById(c *gin.Context, id uuid.UUID)
	Find(c *gin.Context, filter any)
	FindAll(c *gin.Context)
	Delete(c *gin.Context, id uuid.UUID)
}

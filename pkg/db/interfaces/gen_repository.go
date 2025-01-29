package interfaces

import (
	"cameron.io/gin-server/pkg/db/data"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenRepository interface {
	Insert(c *gin.Context, entity any) error
	Upsert(c *gin.Context, filter map[string]any, entity any) (data.Obj, error)
	FindById(c *gin.Context, id uuid.UUID) (data.Obj, error)
	Find(c *gin.Context, filter map[string]any) (data.Obj, error)
	FindAll(c *gin.Context, limit int) ([]data.Obj, error)
	Delete(c *gin.Context, filter map[string]any) (bool, error)
}

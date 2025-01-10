package interfaces

import (
	"cameron.io/gin-server/infra/data"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenRepository interface {
	Insert(c *gin.Context, entity interface{}) error
	Upsert(c *gin.Context, filter map[string]any, entity interface{}) (data.Obj, error)
	FindById(c *gin.Context, id uuid.UUID) (data.Obj, error)
	Find(c *gin.Context, filter map[string]any) (data.Obj, error)
	FindAll(c *gin.Context, limit int) ([]data.Obj, error)
	Delete(c *gin.Context, filter map[string]any) (bool, error)
}

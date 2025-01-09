package repositories

import (
	"cameron.io/gin-server/infra/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenRepository interface {
	Insert(c *gin.Context, entity db.Obj) error
	Upsert(c *gin.Context, filter map[string]interface{}, entity db.Obj) (db.Obj, error)
	FindById(c *gin.Context, id uuid.UUID) (db.Obj, error)
	Find(c *gin.Context, filter map[string]interface{}) (db.Obj, error)
	FindAll(c *gin.Context, limit int) ([]db.Obj, error)
	Delete(c *gin.Context, id uuid.UUID) (bool, error)
}

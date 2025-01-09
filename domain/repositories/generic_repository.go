package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenericRepository interface {
	Insert(c *gin.Context, entity interface{}) error
	Upsert(c *gin.Context, entity interface{}) error
	FindById(c *gin.Context, id uuid.UUID) (interface{}, error)
	Find(c *gin.Context, filter any) (interface{}, error)
	FindAll(c *gin.Context) (interface{}, error)
	Delete(c *gin.Context, id uuid.UUID) error
}

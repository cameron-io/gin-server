package interfaces

import (
	"cameron.io/gin-server/domain/entities"
	"cameron.io/gin-server/domain/utils/data"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	FindUserByEmail(c *gin.Context, email string) (data.Obj, error)
	CreateUser(c *gin.Context, new_user entities.User) error
	DeleteUserByID(c *gin.Context, userId string) (bool, error)
}

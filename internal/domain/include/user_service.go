package include

import (
	"cameron.io/gin-server/internal/domain/models"
	"cameron.io/gin-server/pkg/utils/data"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	FindUserByEmail(c *gin.Context, email string) (data.Obj, error)
	CreateUser(c *gin.Context, new_user models.User) error
	DeleteUserByID(c *gin.Context, userId string) (bool, error)
}

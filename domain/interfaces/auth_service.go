package interfaces

import (
	"cameron.io/gin-server/api/dto"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authenticator(ctx *gin.Context, userEmail string, userPassword string) (*dto.Identity, error)
}

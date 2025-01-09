package interfaces

import "github.com/gin-gonic/gin"

type AuthService interface {
	Authenticator() func(c *gin.Context) (interface{}, error)
}

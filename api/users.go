package api

import (
	"net/http"
	"time"

	"cameron.io/gin-server/models"
	"cameron.io/gin-server/services"
	"cameron.io/gin-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserCreate(ctx *gin.Context, new_user models.User) (*mongo.InsertOneResult, error) {

	// Check if user already exists
	existing_user, err := services.FindUserByEmail(ctx, new_user.Email)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
		return nil, err
	}
	if existing_user != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "user already exists"})
		return nil, err
	}

	// Hash password
	if new_user.Password, err = utils.HashPassword(new_user.Password); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"unexpected_error": err.Error()})
		return nil, err
	}

	// Create new user
	new_user.CreatedAt = time.Now().UnixMilli()

	return services.CreateUser(ctx, new_user)
}

func UserAuth(ctx *gin.Context, user_auth models.Auth) (string, error) {

	existing_user, err := services.FindUserByEmail(ctx, user_auth.Email)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
		return "", err
	}
	if existing_user == nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"msg": "invalid credentials."})
		return "", err
	}
	is_match := utils.CheckPasswordHash(user_auth.Password)
	if !is_match {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "invalid credentials."})
		return "", err
	}
	return utils.CreateToken(user_auth.Email)
}

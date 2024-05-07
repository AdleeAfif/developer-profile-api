package middlewares

import (
	"fmt"
	"net/http"
	model "project/developer-profile-api/models"
	"project/developer-profile-api/utils"

	"github.com/gin-gonic/gin"
)

func Authorize(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Action Not Authorized (Admin only!)"})
		return
	}

	email, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := model.GetUserByEmail(email)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	if user.Role != "admin" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("User %v is not an admin", user.Email)})
	}

	context.Next()
}

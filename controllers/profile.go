package controller

import (
	"net/http"
	model "project/developer-profile-api/models"

	"github.com/gin-gonic/gin"
)

func AddProfile(context *gin.Context) {
	var profile model.Profile

	if err := context.ShouldBindJSON(&profile); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := profile.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"result": result})
}

package controller

import (
	"net/http"
	"os"
	model "project/developer-profile-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func ViewProfile(context *gin.Context) {
	profile, err := model.GetLatestProfile()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"profile": profile})
}

func UpdateProfile(context *gin.Context) {
	profileID, err := primitive.ObjectIDFromHex(os.Getenv("DEFAULT_ID"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse default ID"})
		return
	}

	var payload map[string]interface{}

	err = context.ShouldBindJSON(&payload)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	payload["updatedAt"] = time.Now()

	filter := bson.D{{Key: "_id", Value: profileID}}
	update := bson.D{{Key: "$set", Value: payload}}

	result, err := model.UpdateByID(filter, update)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"result": result})
}

package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"project/developer-profile-api/middlewares"
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

	cacheData, _ := middlewares.GetCache(context, middlewares.RedisClient, context.ClientIP(), context.Request.URL.Path)
	if cacheData != "" {
		var dataToJson map[string]interface{}
		json.Unmarshal([]byte(cacheData), &dataToJson)
		context.JSON(http.StatusOK, dataToJson)
		return
	}

	profile, err := model.GetLatestProfile()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	newProfile := map[string]interface{}{
		"Updated At":   profile.UpdatedAt,
		"Bio":          profile.Bio,
		"Location":     profile.Location.State + ", " + profile.Location.Country,
		"Experience":   profile.Work,
		"Skills":       profile.Skills,
		"Certificates": profile.Certificates,
		"Education":    profile.Education,
		"Social Links": profile.SocialLinks,
	}

	respond := gin.H{"About " + profile.Name.FirstName + " " + profile.Name.LastName: newProfile}

	middlewares.SetCache(context, middlewares.RedisClient, context.ClientIP(), context.Request.URL.Path, respond, 60*time.Minute)

	context.JSON(http.StatusOK, respond)
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

func DownloadResume(context *gin.Context) {
	// Path to your PDF file
	pdfPath := "./misc/resume-adlee-18052024.pdf"

	// Get the file name
	fileName := filepath.Base(pdfPath)

	// Set the Content-Disposition header to prompt download
	context.Header("Content-Disposition", "attachment; filename="+fileName)

	// Serve the file
	context.File(pdfPath)
}

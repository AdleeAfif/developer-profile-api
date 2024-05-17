package main

import (
	"fmt"
	"net/http"
	"os"
	"project/developer-profile-api/db"
	router "project/developer-profile-api/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Get Client, Context and
	// err from connect method.
	client, ctx, err := db.Init(os.Getenv("MONGO_URI"))
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer client.Disconnect(ctx)

	server := gin.Default()

	server.Use(cors.Default())

	server.GET("/", getDefault)

	router.ProfileRoutes(server)
	router.UserRoutes(server)

	fmt.Println("Current running version: 1.2.0")

	server.Run(os.Getenv("PORT"))
}

func getDefault(context *gin.Context) {
	routes := []map[string]string{
		{"path": "/", "description": "Default route. Shows available routes."},
		{"path": "/profile", "description": "This endpoint returns a detailed JSON object containing the profile information of Nik Adlee Afif Nik Mohd Kamil. The profile includes biographical details, certifications, education, work experience, location, skills, social links, and the last updated timestamp."},
		// Add more routes as needed
	}

	context.JSON(http.StatusOK, gin.H{"available_routes": routes})
}

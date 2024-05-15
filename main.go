package main

import (
	"net/http"
	"os"
	"project/developer-profile-api/db"
	router "project/developer-profile-api/routers"
	"time"

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

	// CORS configuration
	config := cors.Config{
		AllowAllOrigins:  true,                                     // Replace with your front-end domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	server.Use(cors.New(config))

	server.GET("/", getDefault)

	router.ProfileRoutes(server)
	router.UserRoutes(server)

	server.Run(os.Getenv("PORT"))
}

func getDefault(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This is default route"})
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"project/developer-profile-api/db"
	middlewares "project/developer-profile-api/middlewares"
	router "project/developer-profile-api/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	redisClient := middlewares.InitRedis()
	defer redisClient.Close()

	mongoClient, mongoCtx, err := db.Init()
	if err != nil {
		panic(err)
	}
	defer mongoClient.Disconnect(mongoCtx)

	server := gin.Default()

	server.Use(cors.Default())

	server.GET("/", getDefault)

	router.ProfileRoutes(server)
	router.UserRoutes(server)

	fmt.Println("Current running version: 1.3.0")

	server.Run(os.Getenv("PORT"))
}

func getDefault(context *gin.Context) {
	routes := []map[string]string{
		{"path": "/", "description": "Default route. Shows available routes."},
		{"path": "/profile", "description": "Returns a detailed JSON object containing the profile information of myself!"},
		{"path": "/resume", "description": "Interested in employing me? Get my resume here!"},
		// Add more routes as needed
	}

	context.JSON(http.StatusOK, gin.H{"available_routes": routes})
}

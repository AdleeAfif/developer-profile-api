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

	fmt.Println("Current running version: 1.1.1")

	server.Run(os.Getenv("PORT"))
}

func getDefault(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This is default route"})
}

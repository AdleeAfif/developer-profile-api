package main

import (
	"fmt"
	"net/http"
	"os"
	"project/developer-profile-api/db"
	router "project/developer-profile-api/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Get Client, Context, CancelFunc and
	// err from connect method.
	client, ctx, err := db.Init(os.Getenv("MONGO_URI"))
	server := gin.Default()
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer client.Disconnect(ctx)
	server.GET("/", getDefault)

	router.ProfileRoutes(server)

	fmt.Println(db.ProfileCollection)

	server.Run(os.Getenv("PORT"))
}

func getDefault(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This is default route"})
}

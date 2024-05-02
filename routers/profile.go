package router

import (
	controller "project/developer-profile-api/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(server *gin.Engine) {
	server.GET("/profile", controller.ViewProfile).POST("/profile", controller.AddProfile).PUT("/profile", controller.UpdateProfile)
	server.GET("/resume")
}

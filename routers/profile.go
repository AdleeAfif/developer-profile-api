package router

import (
	controller "project/developer-profile-api/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(server *gin.Engine) {
	server.POST("/profile", controller.AddProfile)
	server.GET("/resume")
}

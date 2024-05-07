package router

import (
	controller "project/developer-profile-api/controllers"
	"project/developer-profile-api/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(server *gin.Engine) {
	server.GET("/profile", controller.ViewProfile).POST("/profile", middlewares.Authorize, controller.AddProfile).PUT("/profile", middlewares.Authorize, controller.UpdateProfile)
	server.GET("/resume")
}

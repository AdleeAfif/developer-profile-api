package router

import (
	controller "project/developer-profile-api/controllers"
	"project/developer-profile-api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(server *gin.Engine) {
	server.POST("/add", middlewares.Authorize, controller.AddUser)
	server.POST("/login", controller.Login)
}

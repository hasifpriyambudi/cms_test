package routes

import (
	"github.com/gin-gonic/gin"
	authController "github.com/hasifpriyambudi/cms_test/controllers/auth"
)

func Auth(route *gin.RouterGroup, authController authController.AuthController) {
	routes := route.Group("/auth")
	{
		routes.POST("/login", authController.Login)
	}
}

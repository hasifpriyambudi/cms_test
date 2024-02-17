package routes

import "github.com/gin-gonic/gin"

func Admin(route *gin.RouterGroup) {
	routes := route.Group("/test")
	{
		routes.GET("/test", func(ctx *gin.Context) {

		})
	}
}

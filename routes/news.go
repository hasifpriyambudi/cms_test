package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/controllers"
)

func News(route *gin.RouterGroup, newsController controllers.NewsController) {
	routes := route.Group("/news")
	{
		routes.GET("/get", newsController.GetNews)
		routes.GET("/:id", newsController.DetailNews)
	}
}

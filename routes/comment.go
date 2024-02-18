package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/controllers"
)

func Comment(route *gin.RouterGroup, commentController controllers.CommentController) {
	routes := route.Group("/comment")
	{
		routes.POST("/add", commentController.CreateCommentNews)
	}
}

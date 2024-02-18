package routesadmin

import (
	"github.com/gin-gonic/gin"
	newsadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/news"
)

func News(route *gin.RouterGroup, newsController newsadmincontroller.NewsAdminController) {
	routes := route.Group("/news")
	{
		routes.POST("/add", newsController.CreateNewsAdmin)
		// routes.GET("/get", newsController.GetNewsAdmin)
		// routes.GET("/:id", newsController.DetailCustomPageAdmin)
		routes.DELETE("/delete/:id", newsController.DeleteNewsAdmin)
		routes.PUT("/update", newsController.UpdateNewsAdmin)
	}
}

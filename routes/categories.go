package routes

import (
	"github.com/gin-gonic/gin"
	categoriesadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/categories"
)

func Categories(route *gin.RouterGroup, categoriesAdminController categoriesadmincontroller.CategoriesAdminController) {
	routes := route.Group("/categories")
	{
		routes.POST("/add", categoriesAdminController.CreateCategoriesAdmin)
		routes.GET("/get", categoriesAdminController.GetCategoriesAdmin)
		routes.GET("/:id", categoriesAdminController.DetailCategoriesAdmin)
		routes.DELETE("/delete/:id", categoriesAdminController.DeleteCategoriesAdmin)
		routes.PUT("/update", categoriesAdminController.UpdateCategoriesAdmin)
	}
}

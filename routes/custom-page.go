package routes

import (
	"github.com/gin-gonic/gin"
	custompageadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/custom-page"
)

func CustomPage(route *gin.RouterGroup, customPageAdminController custompageadmincontroller.CustomPageAdminController) {
	routes := route.Group("/custom-page")
	{
		routes.POST("/add", customPageAdminController.CreateCustomPageAdmin)
		routes.GET("/get", customPageAdminController.GetCustomPageAdmin)
		routes.GET("/:id", customPageAdminController.DetailCustomPageAdmin)
		routes.DELETE("/delete/:id", customPageAdminController.DeleteCustomPageAdmin)
		routes.PUT("/update", customPageAdminController.UpdateCustomPageAdmin)
	}
}

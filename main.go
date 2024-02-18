package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/app"
	categoriesadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/categories"
	custompageadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/custom-page"
	authController "github.com/hasifpriyambudi/cms_test/controllers/auth"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/middleware"
	categoriesadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/categories"
	custompageadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/custom-page"
	authrepository "github.com/hasifpriyambudi/cms_test/repository/auth"
	"github.com/hasifpriyambudi/cms_test/routes"
	categoriesadminservice "github.com/hasifpriyambudi/cms_test/service/admin/categories"
	custompageadminservice "github.com/hasifpriyambudi/cms_test/service/admin/custom-page"
	authservice "github.com/hasifpriyambudi/cms_test/service/auth"
)

func main() {

	var (
		db       *sql.DB             = app.NewDB()
		validate *validator.Validate = validator.New()

		authRepository authrepository.AuthRepository = authrepository.NewAuthReposisotryImpl()
		authService    authservice.AuthService       = authservice.NewAuthServiceImpl(authRepository, db, validate)
		authController authController.AuthController = authController.NewAuthControllerImpl(authService)

		categoriesRepository categoriesadminrepository.CategoriesAdminRepository = categoriesadminrepository.NewCategoriesAdminRepositoryImpl()
		categoriesService    categoriesadminservice.CategoriesAdminService       = categoriesadminservice.NewCategoriesAdminServiceImpl(categoriesRepository, db, validate)
		categoriesController categoriesadmincontroller.CategoriesAdminController = categoriesadmincontroller.NewCategoriesAdminRepositoryImpl(categoriesService)

		customPageRepository custompageadminrepository.CustomPageAdminRepository = custompageadminrepository.NewCustomPageAdminRepositoryImpl()
		customPageService    custompageadminservice.CustomPageAdminService       = custompageadminservice.NewCustomPageAdminServiceImpl(customPageRepository, db, validate)
		customPageController custompageadmincontroller.CustomPageAdminController = custompageadmincontroller.NewCustomPageAdminRepositoryImpl(customPageService)
	)

	server := gin.Default()
	apiGroup := server.Group("/api")
	apiGroup.Use(exceptions.ErrorHandler())

	// Auth
	routes.Auth(apiGroup, authController)

	// Admin
	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middleware.Authenticate())
	routes.Categories(adminGroup, categoriesController)
	routes.CustomPage(adminGroup, customPageController)

	// Check Port Running
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Error Running server : %v", err)
	}
}

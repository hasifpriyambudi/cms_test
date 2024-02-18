package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/app"
	"github.com/hasifpriyambudi/cms_test/controllers"
	categoriesadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/categories"
	custompageadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/custom-page"
	newsadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/news"
	authController "github.com/hasifpriyambudi/cms_test/controllers/auth"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/middleware"
	categoriesadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/categories"
	custompageadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/custom-page"
	newsadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/news"
	authrepository "github.com/hasifpriyambudi/cms_test/repository/auth"
	"github.com/hasifpriyambudi/cms_test/routes"
	routesadmin "github.com/hasifpriyambudi/cms_test/routes/admin"
	services "github.com/hasifpriyambudi/cms_test/service"
	categoriesadminservice "github.com/hasifpriyambudi/cms_test/service/admin/categories"
	custompageadminservice "github.com/hasifpriyambudi/cms_test/service/admin/custom-page"
	newsadminservice "github.com/hasifpriyambudi/cms_test/service/admin/news"
	authservice "github.com/hasifpriyambudi/cms_test/service/auth"
)

func main() {

	var (
		db       *sql.DB             = app.NewDB()
		validate *validator.Validate = validator.New()

		authRepository authrepository.AuthRepository = authrepository.NewAuthReposisotryImpl()
		authService    authservice.AuthService       = authservice.NewAuthServiceImpl(authRepository, db, validate)
		authController authController.AuthController = authController.NewAuthControllerImpl(authService)

		categoriesRepository      categoriesadminrepository.CategoriesAdminRepository = categoriesadminrepository.NewCategoriesAdminRepositoryImpl()
		categoriesAdminService    categoriesadminservice.CategoriesAdminService       = categoriesadminservice.NewCategoriesAdminServiceImpl(categoriesRepository, db, validate)
		categoriesAdminController categoriesadmincontroller.CategoriesAdminController = categoriesadmincontroller.NewCategoriesAdminRepositoryImpl(categoriesAdminService)

		customPageRepository      custompageadminrepository.CustomPageAdminRepository = custompageadminrepository.NewCustomPageAdminRepositoryImpl()
		customPageAdminService    custompageadminservice.CustomPageAdminService       = custompageadminservice.NewCustomPageAdminServiceImpl(customPageRepository, db, validate)
		customPageAdminController custompageadmincontroller.CustomPageAdminController = custompageadmincontroller.NewCustomPageAdminRepositoryImpl(customPageAdminService)

		newsRepository      newsadminrepository.NewsAdminRepository = newsadminrepository.NewNewsAdminRepositoryImpl()
		newsAdminService    newsadminservice.NewsAdminService       = newsadminservice.NewNewsAdminRepositoryImpl(newsRepository, categoriesRepository, db, validate)
		newsAdminController newsadmincontroller.NewsAdminController = newsadmincontroller.NewNewsAdminRepositoryImpl(newsAdminService)

		newsService    services.NewsService       = services.NewNewsAdminRepositoryImpl(newsRepository, categoriesRepository, db, validate)
		newsController controllers.NewsController = controllers.NewNewsRepositoryImpl(newsService)
	)

	server := gin.Default()
	apiGroup := server.Group("/api")
	apiGroup.Use(exceptions.ErrorHandler())

	// Public
	routes.News(apiGroup, newsController)

	// Auth
	routesadmin.Auth(apiGroup, authController)

	// Admin
	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middleware.Authenticate())
	routesadmin.Categories(adminGroup, categoriesAdminController)
	routesadmin.CustomPage(adminGroup, customPageAdminController)
	routesadmin.News(adminGroup, newsAdminController)

	// Check Port Running
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Error Running server : %v", err)
	}
}

package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/app"
	authController "github.com/hasifpriyambudi/cms_test/controllers/auth"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/middleware"
	authrepository "github.com/hasifpriyambudi/cms_test/repository/auth"
	"github.com/hasifpriyambudi/cms_test/routes"
	authservice "github.com/hasifpriyambudi/cms_test/service/auth"
)

func main() {

	var (
		db             *sql.DB                       = app.NewDB()
		validate       *validator.Validate           = validator.New()
		authRepository authrepository.AuthRepository = authrepository.NewAuthReposisotryImpl()
		authService    authservice.AuthService       = authservice.NewAuthServiceImpl(authRepository, db, validate)
		authController authController.AuthController = authController.NewAuthController(authService)
	)

	server := gin.Default()
	apiGroup := server.Group("/api")
	apiGroup.Use(exceptions.ErrorHandler())

	// Auth
	routes.Auth(apiGroup, authController)

	// Admin
	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middleware.Authenticate())
	routes.Admin(adminGroup)

	// Check Port Running
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Error Running server : %v", err)
	}
}

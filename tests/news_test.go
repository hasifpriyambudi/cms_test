package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/app"
	"github.com/hasifpriyambudi/cms_test/controllers"
	categoriesadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/categories"
	newsadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/news"
	services "github.com/hasifpriyambudi/cms_test/service"
	"github.com/stretchr/testify/assert"
)

func SetupControllerNews() controllers.NewsController {
	var (
		db             = app.NewDBTest()
		validate       = validator.New()
		newsRepo       = newsadminrepository.NewNewsAdminRepositoryImpl()
		categoryRepo   = categoriesadminrepository.NewCategoriesAdminRepositoryImpl()
		newsService    = services.NewNewsAdminRepositoryImpl(newsRepo, categoryRepo, db, validate)
		newsController = controllers.NewNewsControllerImpl(newsService)
	)
	return newsController
}

func TestGetNews(t *testing.T) {
	server := SetupRoutes()
	newsController := SetupControllerNews()
	server.GET("/api/news/get", newsController.GetNews)

	req, _ := http.NewRequest("GET", "/api/news/get", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDetailNews(t *testing.T) {
	server := SetupRoutes()
	newsController := SetupControllerNews()
	server.GET("/api/news/:id", newsController.GetNews)

	req, _ := http.NewRequest("GET", "/api/news/1", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

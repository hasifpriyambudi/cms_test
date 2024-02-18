package testsadmin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/app"
	newsadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/news"
	"github.com/hasifpriyambudi/cms_test/entity"
	categoriesadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/categories"
	newsadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/news"
	newsadminservice "github.com/hasifpriyambudi/cms_test/service/admin/news"
	"github.com/stretchr/testify/assert"
)

func SetupControllerNews() newsadmincontroller.NewsAdminController {
	var (
		db             = app.NewDBTest()
		newsRepo       = newsadminrepository.NewNewsAdminRepositoryImpl()
		categoriesRepo = categoriesadminrepository.NewCategoriesAdminRepositoryImpl()
		validate       = validator.New()
		newsService    = newsadminservice.NewNewsAdminServiceImpl(newsRepo, categoriesRepo, db, validate)
		newsController = newsadmincontroller.NewNewsAdminControllerImpl(newsService)
	)
	return newsController
}

func TestNewsCreate(t *testing.T) {
	server := SetupRoutes()
	newsController := SetupControllerNews()
	server.POST("/api/admin/news/add", newsController.CreateNewsAdmin)
	customPage := entity.NewsEntity{
		Category_Id:  3,
		News_Content: "Ini News Content",
	}

	jsonValue, _ := json.Marshal(customPage)
	req, _ := http.NewRequest("POST", "/api/admin/news/add", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewsUpdate(t *testing.T) {
	server := SetupRoutes()
	newsController := SetupControllerNews()
	server.PUT("/api/admin/news/update", newsController.UpdateNewsAdmin)
	customPage := entity.NewsEntity{
		Id:           3,
		Category_Id:  4,
		News_Content: "Ini Page Content Update",
	}

	jsonValue, _ := json.Marshal(customPage)
	req, _ := http.NewRequest("PUT", "/api/admin/news/update", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewsDelete(t *testing.T) {
	server := SetupRoutes()
	newsController := SetupControllerNews()
	server.DELETE("/api/admin/news/delete/:id", newsController.DeleteNewsAdmin)

	req, _ := http.NewRequest("DELETE", "/api/admin/news/delete/1", nil)

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

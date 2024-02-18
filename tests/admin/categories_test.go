package testsadmin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/app"
	categoriesadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/categories"
	"github.com/hasifpriyambudi/cms_test/entity"
	categoriesadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/categories"
	categoriesadminservice "github.com/hasifpriyambudi/cms_test/service/admin/categories"
	"github.com/stretchr/testify/assert"
)

func SetupControllerCategories() categoriesadmincontroller.CategoriesAdminController {
	var (
		db                   = app.NewDBTest()
		categoriesRepo       = categoriesadminrepository.NewCategoriesAdminRepositoryImpl()
		validate             = validator.New()
		categoriesService    = categoriesadminservice.NewCategoriesAdminServiceImpl(categoriesRepo, db, validate)
		categoriesController = categoriesadmincontroller.NewCategoriesAdminRepositoryImpl(categoriesService)
	)
	return categoriesController
}

func TestCategoriesCreate(t *testing.T) {
	server := SetupRoutes()
	categoriesController := SetupControllerCategories()
	server.POST("/api/admin/categories/add", categoriesController.CreateCategoriesAdmin)
	categories := entity.CategoriesEntity{
		Name: "Category Test 1",
	}

	jsonValue, _ := json.Marshal(categories)
	req, _ := http.NewRequest("POST", "/api/admin/categories/add", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoriesUpdate(t *testing.T) {
	server := SetupRoutes()
	categoriesController := SetupControllerCategories()
	server.PUT("/api/admin/categories/update", categoriesController.UpdateCategoriesAdmin)
	categories := entity.CategoriesEntity{
		Id:   4,
		Name: "Category Test 1",
	}

	jsonValue, _ := json.Marshal(categories)
	req, _ := http.NewRequest("PUT", "/api/admin/categories/update", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoriesDelete(t *testing.T) {
	server := SetupRoutes()
	categoriesController := SetupControllerCategories()
	server.DELETE("/api/admin/categories/delete/:id", categoriesController.DeleteCategoriesAdmin)

	req, _ := http.NewRequest("DELETE", "/api/admin/categories/delete/2", nil)

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoriesGet(t *testing.T) {
	server := SetupRoutes()
	newsController := SetupControllerCategories()
	server.GET("/api/admin/categories/get", newsController.GetCategoriesAdmin)

	req, _ := http.NewRequest("GET", "/api/admin/categories/get", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoriesDetail(t *testing.T) {
	server := SetupRoutes()
	newsController := SetupControllerCategories()
	server.GET("/api/admin/categories/:id", newsController.DetailCategoriesAdmin)

	req, _ := http.NewRequest("GET", "/api/admin/categories/4", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

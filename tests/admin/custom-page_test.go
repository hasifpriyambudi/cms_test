package testsadmin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/app"
	custompageadmincontroller "github.com/hasifpriyambudi/cms_test/controllers/admin/custom-page"
	"github.com/hasifpriyambudi/cms_test/entity"
	custompageadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/custom-page"
	custompageadminservice "github.com/hasifpriyambudi/cms_test/service/admin/custom-page"
	"github.com/stretchr/testify/assert"
)

func SetupControllerCustomPage() custompageadmincontroller.CustomPageAdminController {
	var (
		db                   = app.NewDBTest()
		customPageRepo       = custompageadminrepository.NewCustomPageAdminRepositoryImpl()
		validate             = validator.New()
		customPageService    = custompageadminservice.NewCustomPageAdminServiceImpl(customPageRepo, db, validate)
		customPageController = custompageadmincontroller.NewCustomPageAdminRepositoryImpl(customPageService)
	)
	return customPageController
}

func TestCustomPageCreate(t *testing.T) {
	server := SetupRoutes()
	customPageController := SetupControllerCustomPage()
	server.POST("/api/admin/custom-page/add", customPageController.CreateCustomPageAdmin)
	customPage := entity.CustomPageEntity{
		Custom_Url:   "/asdasd-sadad/",
		Page_Content: "Ini Page Content",
	}

	jsonValue, _ := json.Marshal(customPage)
	req, _ := http.NewRequest("POST", "/api/admin/custom-page/add", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCustomPageUpdate(t *testing.T) {
	server := SetupRoutes()
	customPageController := SetupControllerCustomPage()
	server.PUT("/api/admin/custom-page/update", customPageController.UpdateCustomPageAdmin)
	customPage := entity.CustomPageEntity{
		Id:           8,
		Custom_Url:   "/asdsfsddasd-sadad/",
		Page_Content: "Ini Page Content",
	}

	jsonValue, _ := json.Marshal(customPage)
	req, _ := http.NewRequest("PUT", "/api/admin/custom-page/update", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCustomPageDelete(t *testing.T) {
	server := SetupRoutes()
	customPageController := SetupControllerCustomPage()
	server.DELETE("/api/admin/custom-page/delete/:id", customPageController.DeleteCustomPageAdmin)

	req, _ := http.NewRequest("DELETE", "/api/admin/custom-page/delete/8", nil)

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCustomPageGet(t *testing.T) {
	server := SetupRoutes()
	customPageController := SetupControllerCustomPage()
	server.GET("/api/admin/custom-page/get", customPageController.GetCustomPageAdmin)

	req, _ := http.NewRequest("GET", "/api/admin/custom-page/get", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCustomPageDetail(t *testing.T) {
	server := SetupRoutes()
	customPageController := SetupControllerCustomPage()
	server.GET("/api/admin/custom-page/:id", customPageController.DetailCustomPageAdmin)

	req, _ := http.NewRequest("GET", "/api/admin/custom-page/9", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

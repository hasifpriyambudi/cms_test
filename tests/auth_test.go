package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/app"
	authController "github.com/hasifpriyambudi/cms_test/controllers/auth"
	"github.com/hasifpriyambudi/cms_test/entity"
	authrepository "github.com/hasifpriyambudi/cms_test/repository/auth"
	authservice "github.com/hasifpriyambudi/cms_test/service/auth"
	"github.com/stretchr/testify/assert"
)

func SetupControllerAuth() authController.AuthController {
	var (
		db             = app.NewDBTest()
		authRepo       = authrepository.NewAuthReposisotryImpl()
		validate       = validator.New()
		authService    = authservice.NewAuthServiceImpl(authRepo, db, validate)
		authController = authController.NewAuthControllerImpl(authService)
	)
	return authController
}

func TestAuthPost(t *testing.T) {
	server := SetupRoutes()
	authController := SetupControllerAuth()
	server.POST("/api/auth/login", authController.Login)
	auth := entity.UserEntity{
		Username: "adminadmin",
		Password: "20.11.3309",
	}

	jsonValue, _ := json.Marshal(auth)
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

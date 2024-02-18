package authController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	authservice "github.com/hasifpriyambudi/cms_test/service/auth"
	"github.com/hasifpriyambudi/cms_test/utils"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type AuthControllerImpl struct {
	AuthService authservice.AuthService
}

func NewAuthControllerImpl(service authservice.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: service,
	}
}

func (impl *AuthControllerImpl) Login(ctx *gin.Context) {
	var auth domain.AuthRequest
	err := helpers.ReadJSON(ctx, &auth)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	authResponse := impl.AuthService.Login(ctx, auth)
	res := utils.BuildResponseSuccess(http.StatusOK, "Berhasil", authResponse)
	ctx.JSON(http.StatusOK, res)
}

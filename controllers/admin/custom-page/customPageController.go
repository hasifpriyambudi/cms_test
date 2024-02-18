package custompageadmincontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	custompageadminservice "github.com/hasifpriyambudi/cms_test/service/admin/custom-page"
	"github.com/hasifpriyambudi/cms_test/utils"
)

type CustomPageAdminController interface {
	CreateCustomPageAdmin(ctx *gin.Context)
	UpdateCustomPageAdmin(ctx *gin.Context)
	DeleteCustomPageAdmin(ctx *gin.Context)
	GetCustomPageAdmin(ctx *gin.Context)
	DetailCustomPageAdmin(ctx *gin.Context)
}

type CustomPageControllerAdminImpl struct {
	CustomPageService custompageadminservice.CustomPageAdminService
}

func NewCustomPageAdminRepositoryImpl(CustomPageService custompageadminservice.CustomPageAdminService) CustomPageAdminController {
	return &CustomPageControllerAdminImpl{
		CustomPageService: CustomPageService,
	}
}

func (impl *CustomPageControllerAdminImpl) CreateCustomPageAdmin(ctx *gin.Context) {
	var customPage domain.CustomPageCreateRequest
	err := helpers.ReadJSON(ctx, &customPage)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	customPageResponse := impl.CustomPageService.CreateCustomPageAdmin(ctx, customPage)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "success", customPageResponse)
	ctx.JSON(http.StatusOK, res)
}

func (impl *CustomPageControllerAdminImpl) UpdateCustomPageAdmin(ctx *gin.Context) {
	var customPages domain.CustomPageUpdateRequest
	err := helpers.ReadJSON(ctx, &customPages)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	customPageResponse := impl.CustomPageService.UpdateCustomPageAdmin(ctx, customPages)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "success", customPageResponse)
	ctx.JSON(http.StatusOK, res)
}

func (impl *CustomPageControllerAdminImpl) DeleteCustomPageAdmin(ctx *gin.Context) {
	// Get Param
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	impl.CustomPageService.DeleteCustomPageAdmin(ctx, id)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "Success Delete", nil)
	ctx.JSON(http.StatusOK, res)
}

func (impl *CustomPageControllerAdminImpl) GetCustomPageAdmin(ctx *gin.Context) {
	// Exec Service
	customPageResponse := impl.CustomPageService.GetCustomPageAdmin(ctx)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "Success Delete", customPageResponse)
	ctx.JSON(http.StatusOK, res)
}

func (impl *CustomPageControllerAdminImpl) DetailCustomPageAdmin(ctx *gin.Context) {
	// Get Param
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	customPageResponse := impl.CustomPageService.DetailCustomPageAdmin(ctx, id)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "Success Delete", customPageResponse)
	ctx.JSON(http.StatusOK, res)
}

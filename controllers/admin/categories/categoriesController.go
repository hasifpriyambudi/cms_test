package categoriesadmincontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	categoriesadminservice "github.com/hasifpriyambudi/cms_test/service/admin/categories"
	"github.com/hasifpriyambudi/cms_test/utils"
)

type CategoriesAdminController interface {
	CreateCategoriesAdmin(ctx *gin.Context)
	UpdateCategoriesAdmin(ctx *gin.Context)
	DeleteCategoriesAdmin(ctx *gin.Context)
	GetCategoriesAdmin(ctx *gin.Context)
	DetailCategoriesAdmin(ctx *gin.Context)
}

type CategoriesControllerImpl struct {
	CategoriesService categoriesadminservice.CategoriesAdminService
}

func NewCategoriesAdminRepositoryImpl(categoriesService categoriesadminservice.CategoriesAdminService) CategoriesAdminController {
	return &CategoriesControllerImpl{
		CategoriesService: categoriesService,
	}
}

func (impl *CategoriesControllerImpl) CreateCategoriesAdmin(ctx *gin.Context) {
	var categories domain.CategoriesCreateRequest
	err := helpers.ReadJSON(ctx, &categories)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	categoriesResponse := impl.CategoriesService.CreateCategoriesAdmin(ctx, categories)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "Success", categoriesResponse)
	ctx.JSON(http.StatusOK, res)
}

func (impl *CategoriesControllerImpl) UpdateCategoriesAdmin(ctx *gin.Context) {
	var categories domain.CategoriesUpdateRequest
	err := helpers.ReadJSON(ctx, &categories)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	categoriesResponse := impl.CategoriesService.UpdateCategoriesAdmin(ctx, categories)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "Success Update", categoriesResponse)
	ctx.JSON(http.StatusOK, res)
}

func (impl *CategoriesControllerImpl) DeleteCategoriesAdmin(ctx *gin.Context) {
	// Get Param
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	impl.CategoriesService.DeleteCategoriesAdmin(ctx, id)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "Success Delete", nil)
	ctx.JSON(http.StatusOK, res)
}

func (impl *CategoriesControllerImpl) GetCategoriesAdmin(ctx *gin.Context) {
	// Exec Service
	categoriesResponse := impl.CategoriesService.GetCategoriesAdmin(ctx)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "Success", categoriesResponse)
	ctx.JSON(http.StatusOK, res)
}

func (impl *CategoriesControllerImpl) DetailCategoriesAdmin(ctx *gin.Context) {
	// Get Param
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	categoriesResponse := impl.CategoriesService.DetailCategoriesAdmin(ctx, id)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "Success", categoriesResponse)
	ctx.JSON(http.StatusOK, res)
}

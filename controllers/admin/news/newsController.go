package newsadmincontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	newsadminservice "github.com/hasifpriyambudi/cms_test/service/admin/news"
	"github.com/hasifpriyambudi/cms_test/utils"
)

type NewsAdminController interface {
	CreateNewsAdmin(ctx *gin.Context)
	UpdateNewsAdmin(ctx *gin.Context)
	DeleteNewsAdmin(ctx *gin.Context)
	GetNewsAdmin(ctx *gin.Context)
	DetailNewsAdmin(ctx *gin.Context)
}

type NewsControllerAdminImpl struct {
	NewsService newsadminservice.NewsAdminService
}

func NewNewsAdminControllerImpl(newsService newsadminservice.NewsAdminService) NewsAdminController {
	return &NewsControllerAdminImpl{
		NewsService: newsService,
	}
}

func (impl *NewsControllerAdminImpl) CreateNewsAdmin(ctx *gin.Context) {
	var news domain.NewsCreateRequest
	err := helpers.ReadJSON(ctx, &news)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	newsResponse := impl.NewsService.CreateNewsAdmin(ctx, news)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "success", newsResponse)
	ctx.JSON(http.StatusOK, res)
}

func (impl *NewsControllerAdminImpl) UpdateNewsAdmin(ctx *gin.Context) {
	var news domain.NewsUpdateRequest
	err := helpers.ReadJSON(ctx, &news)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	newsResponse := impl.NewsService.UpdateNewsAdmin(ctx, news)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "success", newsResponse)
	ctx.JSON(http.StatusOK, res)
}

func (impl *NewsControllerAdminImpl) DeleteNewsAdmin(ctx *gin.Context) {
	// Get Param
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	impl.NewsService.DeleteNewsAdmin(ctx, id)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "success delete", nil)
	ctx.JSON(http.StatusOK, res)
}

func (impl *NewsControllerAdminImpl) GetNewsAdmin(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (impl *NewsControllerAdminImpl) DetailNewsAdmin(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

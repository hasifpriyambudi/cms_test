package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	services "github.com/hasifpriyambudi/cms_test/service"
	"github.com/hasifpriyambudi/cms_test/utils"
)

type NewsController interface {
	GetNews(ctx *gin.Context)
	DetailNews(ctx *gin.Context)
}

type NewsControllerImpl struct {
	NewsService services.NewsService
}

func NewNewsControllerImpl(newsService services.NewsService) NewsController {
	return &NewsControllerImpl{
		NewsService: newsService,
	}
}

func (impl *NewsControllerImpl) GetNews(ctx *gin.Context) {
	// Exec Service
	newsResponse := impl.NewsService.GetNews(ctx)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "success", newsResponse)
	ctx.JSON(http.StatusOK, res)
}

func (impl *NewsControllerImpl) DetailNews(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	newsResponse := impl.NewsService.DetailNews(ctx, id)

	// Return
	res := utils.BuildResponseSuccess(http.StatusOK, "success", newsResponse)
	ctx.JSON(http.StatusOK, res)
}

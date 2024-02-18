package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	services "github.com/hasifpriyambudi/cms_test/service"
	"github.com/hasifpriyambudi/cms_test/utils"
)

type CommentController interface {
	CreateCommentNews(ctx *gin.Context)
}

type CommentControllerImpl struct {
	CommentService services.CommentService
}

func NewCommentControllerImpl(commentService services.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (impl *CommentControllerImpl) CreateCommentNews(ctx *gin.Context) {
	var comment domain.CommentCreateRequest
	err := helpers.ReadJSON(ctx, &comment)
	if err != nil {
		panic(exceptions.NewReadJsonError(err))
	}

	// Exec Service
	impl.CommentService.CreateCommentNews(ctx, comment)

	// return
	res := utils.BuildResponseSuccess(http.StatusOK, "success", nil)
	ctx.JSON(http.StatusOK, res)
}

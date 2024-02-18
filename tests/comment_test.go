package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/app"
	"github.com/hasifpriyambudi/cms_test/controllers"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/repository"
	newsadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/news"
	services "github.com/hasifpriyambudi/cms_test/service"
	"github.com/stretchr/testify/assert"
)

func SetupControllerComment() controllers.CommentController {
	var (
		db                = app.NewDBTest()
		commentRepo       = repository.NewCommentRepositoryImpl()
		newsRepo          = newsadminrepository.NewNewsAdminRepositoryImpl()
		validate          = validator.New()
		commentService    = services.NewCommentServiceImpl(commentRepo, newsRepo, db, validate)
		commentController = controllers.NewCommentControllerImpl(commentService)
	)

	return commentController
}

func TestCommentPost(t *testing.T) {
	server := SetupRoutes()
	commentController := SetupControllerComment()
	server.POST("/api/comment/add", commentController.CreateCommentNews)
	comment := entity.CommentEntity{
		News_Id: 1,
		Name:    "Testing Comment",
		Comment: "Ini Isi Comment Testing",
	}

	jsonValue, _ := json.Marshal(comment)
	req, _ := http.NewRequest("POST", "/api/comment/add", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

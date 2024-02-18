package services

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	"github.com/hasifpriyambudi/cms_test/repository"
	newsadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/news"
)

type CommentService interface {
	CreateCommentNews(ctx *gin.Context, request domain.CommentCreateRequest)
}

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	NewsRepository    newsadminrepository.NewsAdminRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCommentServiceImpl(
	commentRepository repository.CommentRepository,
	newsRepository newsadminrepository.NewsAdminRepository,
	db *sql.DB,
	validate *validator.Validate,
) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		NewsRepository:    newsRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (impl *CommentServiceImpl) CreateCommentNews(ctx *gin.Context, request domain.CommentCreateRequest) {
	// Cek Validator
	err := impl.Validate.Struct(request)
	helpers.PanicError(err)

	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Binding Request
	comment := entity.CommentEntity{
		News_Id: request.News_Id,
		Name:    request.Name,
		Comment: request.Comment,
	}

	// Exec Repo News Detail
	_, err = impl.NewsRepository.DetailNewsAdmin(ctx, tx, comment.News_Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Exec Repo Comment
	impl.CommentRepository.CreateCommentNews(ctx, tx, comment)
}

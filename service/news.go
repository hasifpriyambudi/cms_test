package services

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	categoriesadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/categories"
	newsadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/news"
)

type NewsService interface {
	GetNews(ctx *gin.Context) []domain.NewsGetResponse
	DetailNews(ctx *gin.Context, id int) domain.NewsGetResponse
}

type NewsServiceImpl struct {
	NewsRepository       newsadminrepository.NewsAdminRepository
	CategoriesRepository categoriesadminrepository.CategoriesAdminRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewNewsAdminRepositoryImpl(
	newsRepository newsadminrepository.NewsAdminRepository,
	categoriesRepository categoriesadminrepository.CategoriesAdminRepository,
	db *sql.DB,
	validate *validator.Validate,
) NewsService {
	return &NewsServiceImpl{
		NewsRepository:       newsRepository,
		CategoriesRepository: categoriesRepository,
		DB:                   db,
		Validate:             validate,
	}
}

func (impl *NewsServiceImpl) GetNews(ctx *gin.Context) []domain.NewsGetResponse {
	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo News
	newsBatch, err := impl.NewsRepository.GetNewsAdmin(ctx, tx)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Binding Response
	var resNewsBatch []domain.NewsGetResponse
	for _, news := range newsBatch {
		resNews := domain.NewsGetResponse{
			Id:           news.Id,
			Category_Id:  news.Category_Id,
			News_Content: news.News_Content,
		}
		resNewsBatch = append(resNewsBatch, resNews)
	}

	return resNewsBatch
}

func (impl *NewsServiceImpl) DetailNews(ctx *gin.Context, id int) domain.NewsGetResponse {
	// Cek Validator
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo
	news, err := impl.NewsRepository.DetailNewsAdmin(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	return domain.NewsGetResponse{
		Id:           news.Id,
		Category_Id:  news.Category_Id,
		News_Content: news.News_Content,
	}
}

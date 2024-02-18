package newsadminservice

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	categoriesadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/categories"
	newsadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/news"
)

type NewsAdminService interface {
	CreateNewsAdmin(ctx *gin.Context, request domain.NewsCreateRequest) domain.NewsCreateResponse
	UpdateNewsAdmin(ctx *gin.Context, request domain.NewsUpdateRequest) domain.NewsUpdateResponse
	DeleteNewsAdmin(ctx *gin.Context, id int)
	GetNewsAdmin(ctx *gin.Context) []domain.NewsGetResponse
	DetailNewsAdmin(ctx *gin.Context, id int) domain.NewsGetResponse
}

type NewsAdminServiceImpl struct {
	NewsRepository       newsadminrepository.NewsAdminRepository
	CategoriesRepository categoriesadminrepository.CategoriesAdminRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewNewsAdminServiceImpl(newsRepository newsadminrepository.NewsAdminRepository, categoriesRepository categoriesadminrepository.CategoriesAdminRepository, db *sql.DB, validate *validator.Validate) NewsAdminService {
	return &NewsAdminServiceImpl{
		NewsRepository:       newsRepository,
		CategoriesRepository: categoriesRepository,
		DB:                   db,
		Validate:             validate,
	}
}

func (impl *NewsAdminServiceImpl) CreateNewsAdmin(ctx *gin.Context, request domain.NewsCreateRequest) domain.NewsCreateResponse {
	// Cek Validator
	err := impl.Validate.Struct(request)
	helpers.PanicError(err)

	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Binding Request
	news := entity.NewsEntity{
		Category_Id:  request.Category_Id,
		News_Content: request.News_Content,
	}

	// Exec Repo Detail Categories
	_, err = impl.CategoriesRepository.DetailCategoriesAdmin(ctx, tx, request.Category_Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Exec Repo
	news = impl.NewsRepository.CreateNewsAdmin(ctx, tx, news)

	return domain.NewsCreateResponse{
		Id: news.Id,
	}
}

func (impl *NewsAdminServiceImpl) UpdateNewsAdmin(ctx *gin.Context, request domain.NewsUpdateRequest) domain.NewsUpdateResponse {
	// Cek Validator
	err := impl.Validate.Struct(request)
	helpers.PanicError(err)

	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo Detail
	news, err := impl.NewsRepository.DetailNewsAdmin(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Exec Repo Detail Categories
	_, err = impl.CategoriesRepository.DetailCategoriesAdmin(ctx, tx, request.Category_Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Update Binding
	news.News_Content = request.News_Content
	news.Category_Id = request.Category_Id

	// Exec Repo Update
	news = impl.NewsRepository.UpdateNewsAdmin(ctx, tx, news)

	// Return
	return domain.NewsUpdateResponse{
		Id:           news.Id,
		Category_Id:  news.Category_Id,
		News_Content: news.News_Content,
	}

}

func (impl *NewsAdminServiceImpl) DeleteNewsAdmin(ctx *gin.Context, id int) {
	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo Detail
	news, err := impl.NewsRepository.DetailNewsAdmin(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	impl.NewsRepository.DeleteNewsAdmin(ctx, tx, news.Id)
}

func (impl *NewsAdminServiceImpl) GetNewsAdmin(ctx *gin.Context) []domain.NewsGetResponse {
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

func (impl *NewsAdminServiceImpl) DetailNewsAdmin(ctx *gin.Context, id int) domain.NewsGetResponse {
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

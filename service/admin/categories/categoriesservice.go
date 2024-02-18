package categoriesadminservice

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	categoriesadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/categories"
)

type CategoriesAdminService interface {
	CreateCategoriesAdmin(ctx *gin.Context, request domain.CategoriesCreateRequest) domain.CategoriesCreateResponse
	UpdateCategoriesAdmin(ctx *gin.Context, request domain.CategoriesUpdateRequest) domain.CategoriesUpdateResponse
	DeleteCategoriesAdmin(ctx *gin.Context, id int)
	GetCategoriesAdmin(ctx *gin.Context) []domain.CategoriesGetResponse
	DetailCategoriesAdmin(ctx *gin.Context, id int) domain.CategoriesGetResponse
}

type CategoriesAdminServiceImpl struct {
	CategoriesRepository categoriesadminrepository.CategoriesAdminRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewCategoriesAdminServiceImpl(categoriesRepository categoriesadminrepository.CategoriesAdminRepository, db *sql.DB, validate *validator.Validate) CategoriesAdminService {
	return &CategoriesAdminServiceImpl{
		CategoriesRepository: categoriesRepository,
		DB:                   db,
		Validate:             validate,
	}
}

func (impl *CategoriesAdminServiceImpl) CreateCategoriesAdmin(ctx *gin.Context, request domain.CategoriesCreateRequest) domain.CategoriesCreateResponse {
	// Cek Validator
	err := impl.Validate.Struct(request)
	helpers.PanicError(err)

	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Binding Request
	categories := entity.CategoriesEntity{
		Name_Category: request.Name,
	}

	categories = impl.CategoriesRepository.CreateCategoriesAdmin(ctx, tx, categories)

	return domain.CategoriesCreateResponse{
		Id: categories.Id,
	}
}

func (impl *CategoriesAdminServiceImpl) UpdateCategoriesAdmin(ctx *gin.Context, request domain.CategoriesUpdateRequest) domain.CategoriesUpdateResponse {
	// Cek Validator
	err := impl.Validate.Struct(request)
	helpers.PanicError(err)

	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo Detail
	categories, err := impl.CategoriesRepository.DetailCategoriesAdmin(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Update Binding
	categories.Name_Category = request.Name

	// Exec Repo Update
	categories = impl.CategoriesRepository.UpdateCategoriesAdmin(ctx, tx, categories)

	return domain.CategoriesUpdateResponse{
		Id:   categories.Id,
		Name: categories.Name_Category,
	}
}

func (impl *CategoriesAdminServiceImpl) DeleteCategoriesAdmin(ctx *gin.Context, id int) {
	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo Detail
	categories, err := impl.CategoriesRepository.DetailCategoriesAdmin(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Exec Repo Delete
	impl.CategoriesRepository.DeleteCategoriesAdmin(ctx, tx, categories.Id)
}

func (impl *CategoriesAdminServiceImpl) GetCategoriesAdmin(ctx *gin.Context) []domain.CategoriesGetResponse {
	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo Categories
	categories, err := impl.CategoriesRepository.GetCategoriesAdmin(ctx, tx)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Binding Response
	var resCategories []domain.CategoriesGetResponse
	for _, category := range categories {
		resCategory := domain.CategoriesGetResponse{
			Id:   category.Id,
			Name: category.Name_Category,
		}
		resCategories = append(resCategories, resCategory)
	}

	return resCategories
}

func (impl *CategoriesAdminServiceImpl) DetailCategoriesAdmin(ctx *gin.Context, id int) domain.CategoriesGetResponse {
	// Cek Validator
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo
	category, err := impl.CategoriesRepository.DetailCategoriesAdmin(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	return domain.CategoriesGetResponse{
		Id:   category.Id,
		Name: category.Name_Category,
	}
}

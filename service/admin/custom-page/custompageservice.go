package custompageadminservice

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	custompageadminrepository "github.com/hasifpriyambudi/cms_test/repository/admin/custom-page"
)

type CustomPageAdminService interface {
	CreateCustomPageAdmin(ctx *gin.Context, request domain.CustomPageCreateRequest) domain.CustomPageCreateResponse
	UpdateCustomPageAdmin(ctx *gin.Context, request domain.CustomPageUpdateRequest) domain.CustomPageUpdateResponse
	DeleteCustomPageAdmin(ctx *gin.Context, id int)
	GetCustomPageAdmin(ctx *gin.Context) []domain.CustomPageGetResponse
	DetailCustomPageAdmin(ctx *gin.Context, id int) domain.CustomPageGetResponse
}

type CustomPageAdminServiceImpl struct {
	CustomPageRepository custompageadminrepository.CustomPageAdminRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewCustomPageAdminServiceImpl(customPageRepository custompageadminrepository.CustomPageAdminRepository, db *sql.DB, validate *validator.Validate) CustomPageAdminService {
	return &CustomPageAdminServiceImpl{
		CustomPageRepository: customPageRepository,
		DB:                   db,
		Validate:             validate,
	}
}

func (impl *CustomPageAdminServiceImpl) CreateCustomPageAdmin(ctx *gin.Context, request domain.CustomPageCreateRequest) domain.CustomPageCreateResponse {
	// Cek Validator
	err := impl.Validate.Struct(request)
	helpers.PanicError(err)

	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Binding Request
	customPage := entity.CustomPageEntity{
		Custom_Url:   request.Custom_Url,
		Page_Content: request.Page_Content,
	}

	// Exec Repo
	customPage = impl.CustomPageRepository.CreateCustomPageAdmin(ctx, tx, customPage)

	return domain.CustomPageCreateResponse{
		Id: customPage.Id,
	}
}

func (impl *CustomPageAdminServiceImpl) UpdateCustomPageAdmin(ctx *gin.Context, request domain.CustomPageUpdateRequest) domain.CustomPageUpdateResponse {
	// Cek Validator
	err := impl.Validate.Struct(request)
	helpers.PanicError(err)

	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo Detail
	customPage, err := impl.CustomPageRepository.DetailCustomPageAdmin(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Update Binding
	customPage.Custom_Url = request.Custom_Url
	customPage.Page_Content = request.Page_Content

	// Exec Repo Update
	customPage = impl.CustomPageRepository.UpdateCustomPageAdmin(ctx, tx, customPage)

	return domain.CustomPageUpdateResponse{
		Id:           customPage.Id,
		Custom_Url:   customPage.Custom_Url,
		Page_Content: customPage.Page_Content,
	}
}

func (impl *CustomPageAdminServiceImpl) DeleteCustomPageAdmin(ctx *gin.Context, id int) {
	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo Detail
	customPage, err := impl.CustomPageRepository.DetailCustomPageAdmin(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Exec Repo Delete
	impl.CustomPageRepository.DeleteCustomPageAdmin(ctx, tx, customPage.Id)
}

func (impl *CustomPageAdminServiceImpl) GetCustomPageAdmin(ctx *gin.Context) []domain.CustomPageGetResponse {
	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo Custom Page
	customPages, err := impl.CustomPageRepository.GetCustomPageAdmin(ctx, tx)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Binding Response
	var resCustomPages []domain.CustomPageGetResponse
	for _, customPage := range customPages {
		resCustomPage := domain.CustomPageGetResponse{
			Id:           customPage.Id,
			Custom_Url:   customPage.Custom_Url,
			Page_Content: customPage.Page_Content,
		}
		resCustomPages = append(resCustomPages, resCustomPage)
	}

	return resCustomPages
}

func (impl *CustomPageAdminServiceImpl) DetailCustomPageAdmin(ctx *gin.Context, id int) domain.CustomPageGetResponse {
	// Cek Validator
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Exec Repo
	customPage, err := impl.CustomPageRepository.DetailCustomPageAdmin(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	return domain.CustomPageGetResponse{
		Id:           customPage.Id,
		Custom_Url:   customPage.Custom_Url,
		Page_Content: customPage.Page_Content,
	}
}

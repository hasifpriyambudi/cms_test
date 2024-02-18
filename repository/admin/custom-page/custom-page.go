package custompageadminrepository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
)

type CustomPageAdminRepository interface {
	CreateCustomPageAdmin(ctx *gin.Context, tx *sql.Tx, CustomPage entity.CustomPageEntity) entity.CustomPageEntity
	UpdateCustomPageAdmin(ctx *gin.Context, tx *sql.Tx, CustomPage entity.CustomPageEntity) entity.CustomPageEntity
	DeleteCustomPageAdmin(ctx *gin.Context, tx *sql.Tx, id int)
	GetCustomPageAdmin(ctx *gin.Context, tx *sql.Tx) ([]entity.CustomPageEntity, error)
	DetailCustomPageAdmin(ctx *gin.Context, tx *sql.Tx, id int) (entity.CustomPageEntity, error)
}

type CustomPageAdminRepositoryImpl struct{}

func NewCustomPageAdminRepositoryImpl() CustomPageAdminRepository {
	return &CustomPageAdminRepositoryImpl{}
}

func (impl *CustomPageAdminRepositoryImpl) CreateCustomPageAdmin(ctx *gin.Context, tx *sql.Tx, customPage entity.CustomPageEntity) entity.CustomPageEntity {
	sqlQuery := "INSERT INTO custom_page(custom_url, page_content, created_at) VALUE(?, ?, ?)"
	res, err := tx.ExecContext(ctx, sqlQuery, customPage.Custom_Url, customPage.Page_Content, time.Now())
	if err != nil {
		err = helpers.MysqlError(err)
		panic(exceptions.NewMysqlError(err))
	}

	// Get Last Insert
	id, err := res.LastInsertId()
	helpers.PanicError(err)

	// Return
	customPage.Id = int(id)
	return customPage
}

func (impl *CustomPageAdminRepositoryImpl) UpdateCustomPageAdmin(ctx *gin.Context, tx *sql.Tx, customPage entity.CustomPageEntity) entity.CustomPageEntity {
	sqlQuery := "UPDATE custom_page SET custom_url=?, page_content=?, updated_at=? WHERE id=?"
	_, err := tx.ExecContext(ctx, sqlQuery, customPage.Custom_Url, customPage.Page_Content, time.Now(), customPage.Id)
	helpers.PanicError(err)

	return customPage
}

func (impl *CustomPageAdminRepositoryImpl) DeleteCustomPageAdmin(ctx *gin.Context, tx *sql.Tx, id int) {
	sqlQUery := "UPDATE custom_page SET deleted_at=? WHERE id=?"
	_, err := tx.ExecContext(ctx, sqlQUery, time.Now(), id)
	helpers.PanicError(err)
}

func (impl *CustomPageAdminRepositoryImpl) GetCustomPageAdmin(ctx *gin.Context, tx *sql.Tx) ([]entity.CustomPageEntity, error) {
	sqlQuery := "SELECT id, custom_url, page_content FROM custom_page WHERE deleted_at IS NULL"
	res, err := tx.QueryContext(ctx, sqlQuery)
	helpers.PanicError(err)
	defer res.Close()

	// Binding
	var customPages []entity.CustomPageEntity
	for res.Next() {
		customPage := entity.CustomPageEntity{}
		err := res.Scan(&customPage.Id, &customPage.Custom_Url, &customPage.Page_Content)
		helpers.PanicError(err)
		customPages = append(customPages, customPage)
	}

	// Check Return Exists
	if len(customPages) > 0 {
		return customPages, nil
	}

	return customPages, errors.New("custom page not found")
}

func (impl *CustomPageAdminRepositoryImpl) DetailCustomPageAdmin(ctx *gin.Context, tx *sql.Tx, id int) (entity.CustomPageEntity, error) {
	sqlQuery := "SELECT id, custom_url, page_content FROM custom_page WHERE id=? AND deleted_at is NULL"
	res, err := tx.QueryContext(ctx, sqlQuery, id)
	helpers.PanicError(err)
	defer res.Close()

	// Binding
	customPage := entity.CustomPageEntity{}
	if res.Next() {
		err := res.Scan(&customPage.Id, &customPage.Custom_Url, &customPage.Page_Content)
		helpers.PanicError(err)
		return customPage, nil
	}
	return customPage, errors.New("custom page not found")
}

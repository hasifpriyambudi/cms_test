package categoriesadminrepository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
)

type CategoriesAdminRepository interface {
	CreateCategoriesAdmin(ctx *gin.Context, tx *sql.Tx, categories entity.CategoriesEntity) entity.CategoriesEntity
	UpdateCategoriesAdmin(ctx *gin.Context, tx *sql.Tx, categories entity.CategoriesEntity) entity.CategoriesEntity
	DeleteCategoriesAdmin(ctx *gin.Context, tx *sql.Tx, id int)
	GetCategoriesAdmin(ctx *gin.Context, tx *sql.Tx) ([]entity.CategoriesEntity, error)
	DetailCategoriesAdmin(ctx *gin.Context, tx *sql.Tx, id int) (entity.CategoriesEntity, error)
}

type CategoriesAdminRepositoryImpl struct{}

func NewCategoriesAdminRepositoryImpl() CategoriesAdminRepository {
	return &CategoriesAdminRepositoryImpl{}
}

func (impl *CategoriesAdminRepositoryImpl) CreateCategoriesAdmin(ctx *gin.Context, tx *sql.Tx, categories entity.CategoriesEntity) entity.CategoriesEntity {
	sqlQuery := "INSERT INTO category(name_category, created_at) VALUE(?, ?)"
	res, err := tx.ExecContext(ctx, sqlQuery, categories.Name_Category, time.Now())
	if err != nil {
		err = helpers.MysqlError(err)
		panic(exceptions.NewMysqlError(err))
	}

	// Get Last Insert ID
	id, err := res.LastInsertId()
	helpers.PanicError(err)

	categories.Id = int(id)
	return categories
}

func (impl *CategoriesAdminRepositoryImpl) UpdateCategoriesAdmin(ctx *gin.Context, tx *sql.Tx, categories entity.CategoriesEntity) entity.CategoriesEntity {
	sqlQuery := "UPDATE category SET name_category=?, updated_at=? WHERE id=?"
	_, err := tx.ExecContext(ctx, sqlQuery, categories.Name_Category, time.Now(), categories.Id)
	helpers.PanicError(err)

	return categories
}

func (impl *CategoriesAdminRepositoryImpl) DeleteCategoriesAdmin(ctx *gin.Context, tx *sql.Tx, id int) {
	sqlQuery := "UPDATE category SET deleted_at=? WHERE id=?"
	_, err := tx.ExecContext(ctx, sqlQuery, time.Now(), id)
	helpers.PanicError(err)
}

func (impl *CategoriesAdminRepositoryImpl) GetCategoriesAdmin(ctx *gin.Context, tx *sql.Tx) ([]entity.CategoriesEntity, error) {
	sqlQuery := "SELECT id, name_category FROM category WHERE deleted_at is NULL ORDER BY id DESC"
	res, err := tx.QueryContext(ctx, sqlQuery)
	helpers.PanicError(err)
	defer res.Close()

	var categories []entity.CategoriesEntity
	for res.Next() {
		category := entity.CategoriesEntity{}
		err := res.Scan(&category.Id, &category.Name_Category)
		helpers.PanicError(err)
		categories = append(categories, category)
	}

	// Check Return Exists
	if len(categories) > 0 {
		return categories, nil
	}
	return categories, errors.New("categories not found")

}

func (impl *CategoriesAdminRepositoryImpl) DetailCategoriesAdmin(ctx *gin.Context, tx *sql.Tx, id int) (entity.CategoriesEntity, error) {
	sqlQuery := "SELECT id, name_category FROM category WHERE deleted_at is NULL AND id=?"
	res, err := tx.QueryContext(ctx, sqlQuery, id)
	helpers.PanicError(err)
	defer res.Close()

	// Binding
	category := entity.CategoriesEntity{}
	if res.Next() {
		err := res.Scan(&category.Id, &category.Name_Category)
		helpers.PanicError(err)
		return category, nil
	}
	return category, errors.New("categories not found")
}

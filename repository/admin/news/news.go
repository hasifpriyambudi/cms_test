package newsadminrepository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
)

type NewsAdminRepository interface {
	CreateNewsAdmin(ctx *gin.Context, tx *sql.Tx, news entity.NewsEntity) entity.NewsEntity
	UpdateNewsAdmin(ctx *gin.Context, tx *sql.Tx, news entity.NewsEntity) entity.NewsEntity
	DeleteNewsAdmin(ctx *gin.Context, tx *sql.Tx, id int)
	GetNewsAdmin(ctx *gin.Context, tx *sql.Tx) ([]entity.NewsEntity, error)
	DetailNewsAdmin(ctx *gin.Context, tx *sql.Tx, id int) (entity.NewsEntity, error)
}

type NewsAdminRepositoryImpl struct{}

func NewNewsAdminRepositoryImpl() NewsAdminRepository {
	return &NewsAdminRepositoryImpl{}
}

func (impl *NewsAdminRepositoryImpl) CreateNewsAdmin(ctx *gin.Context, tx *sql.Tx, news entity.NewsEntity) entity.NewsEntity {
	sqlQuery := "INSERT INTO news(category_id, news_content) VALUE(?, ?)"
	res, err := tx.ExecContext(ctx, sqlQuery, news.Category_Id, news.News_Content)
	if err != nil {
		err = helpers.MysqlError(err)
		panic(exceptions.NewMysqlError(err))
	}

	// Get Last Insert
	id, err := res.LastInsertId()
	helpers.PanicError(err)

	// return
	news.Id = int(id)
	return news
}

func (impl *NewsAdminRepositoryImpl) UpdateNewsAdmin(ctx *gin.Context, tx *sql.Tx, news entity.NewsEntity) entity.NewsEntity {
	sqlQuery := "UPDATE news SET category_id=?, news_content=?, updated_at=? WHERE id=?"
	_, err := tx.ExecContext(ctx, sqlQuery, news.Category_Id, news.News_Content, time.Now(), news.Id)
	helpers.PanicError(err)

	return news
}

func (impl *NewsAdminRepositoryImpl) DeleteNewsAdmin(ctx *gin.Context, tx *sql.Tx, id int) {
	sqlQuery := "UPDATE news SET deleted_at=? WHERE id=?"
	_, err := tx.ExecContext(ctx, sqlQuery, time.Now(), id)
	helpers.PanicError(err)
}

func (impl *NewsAdminRepositoryImpl) GetNewsAdmin(ctx *gin.Context, tx *sql.Tx) ([]entity.NewsEntity, error) {
	sqlQuery := "SELECT id, category_id, news_content FROM news WHERE deleted_at is NULL ORDER BY id DESC"
	res, err := tx.QueryContext(ctx, sqlQuery)
	helpers.PanicError(err)
	defer res.Close()

	// Binding
	var newsBatch []entity.NewsEntity
	for res.Next() {
		news := entity.NewsEntity{}
		err := res.Scan(&news.Id, &news.Category_Id, &news.News_Content)
		helpers.PanicError(err)
		newsBatch = append(newsBatch, news)
	}

	// CHeck Return Exists
	if len(newsBatch) > 0 {
		return newsBatch, nil
	}

	return newsBatch, errors.New("news not found")
}

func (impl *NewsAdminRepositoryImpl) DetailNewsAdmin(ctx *gin.Context, tx *sql.Tx, id int) (entity.NewsEntity, error) {
	sqlQuery := "SELECT id, category_id, news_content FROM news WHERE deleted_at is NULL AND id=?"
	res, err := tx.QueryContext(ctx, sqlQuery, id)
	helpers.PanicError(err)
	defer res.Close()

	// Binding
	news := entity.NewsEntity{}
	if res.Next() {
		err := res.Scan(&news.Id, &news.Category_Id, &news.News_Content)
		helpers.PanicError(err)
		return news, nil
	}

	return news, errors.New("news not found")
}

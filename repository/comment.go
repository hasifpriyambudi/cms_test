package repository

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
)

type CommentRepository interface {
	CreateCommentNews(ctx *gin.Context, tx *sql.Tx, comment entity.CommentEntity) entity.CommentEntity
}

type CommentRepositoryImpl struct{}

func NewCommentRepositoryImpl() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (impl *CommentRepositoryImpl) CreateCommentNews(ctx *gin.Context, tx *sql.Tx, comment entity.CommentEntity) entity.CommentEntity {
	sqlQuery := "INSERT INTO comment(news_id, name, comment, created_at) VALUE(?, ?, ?, ?)"
	res, err := tx.ExecContext(ctx, sqlQuery, comment.News_Id, comment.Name, comment.Comment, time.Now())
	if err != nil {
		err := helpers.MysqlError(err)
		panic(exceptions.NewMysqlError(err))
	}

	// Get Last Insert
	id, err := res.LastInsertId()
	helpers.PanicError(err)

	// return
	comment.Id = int(id)
	return comment
}

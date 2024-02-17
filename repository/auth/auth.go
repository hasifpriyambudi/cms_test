package authrepository

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/helpers"
)

type AuthRepository interface {
	Login(ctx *gin.Context, db *sql.Tx, info entity.UserEntity) (domain.AuthDatabase, error)
}

type AuthRepositoryImpl struct{}

func NewAuthReposisotryImpl() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (impl *AuthRepositoryImpl) Login(ctx *gin.Context, db *sql.Tx, info entity.UserEntity) (domain.AuthDatabase, error) {
	sqlQuery := "SELECT id, name, username, password FROM user WHERE username=? AND deleted_at IS NULL"
	res, err := db.QueryContext(ctx, sqlQuery, info.Username)
	helpers.PanicError(err)
	defer res.Close()

	// Change Rows to Domain
	auth := domain.AuthDatabase{}
	if res.Next() {
		err2 := res.Scan(&auth.Id, &auth.Name, &auth.Username, &auth.Password)
		helpers.PanicError(err2)
		return auth, nil
	}
	return auth, errors.New("account not found")
}

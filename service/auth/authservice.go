package authservice

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/domain"
	"github.com/hasifpriyambudi/cms_test/entity"
	"github.com/hasifpriyambudi/cms_test/exceptions"
	"github.com/hasifpriyambudi/cms_test/helpers"
	authrepository "github.com/hasifpriyambudi/cms_test/repository/auth"
)

type AuthService interface {
	Login(ctx *gin.Context, request domain.AuthRequest) domain.AuthResponse
}

type AuthServiceImpl struct {
	AuthRepository authrepository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthServiceImpl(authRepo authrepository.AuthRepository, db *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepo,
		DB:             db,
		Validate:       validate,
	}
}

func (impl *AuthServiceImpl) Login(ctx *gin.Context, request domain.AuthRequest) domain.AuthResponse {

	// Validator
	err := impl.Validate.Struct(request)
	helpers.PanicError(err)

	// Init Transaction
	tx, err := impl.DB.Begin()
	helpers.PanicError(err)
	defer helpers.CommitOrRollback(tx)

	// Convert To Entity
	authEntity := entity.UserEntity{
		Username: request.Username,
	}

	// Exec Repo
	infoLogin, err := impl.AuthRepository.Login(ctx, tx, authEntity)
	if err != nil {
		panic(exceptions.NewNotFoundError(err))
	}

	// Check Password
	if !helpers.CheckPassword(infoLogin.Password, request.Password) {
		err := errors.New("wrong password")
		panic(exceptions.NewAuthError(err))
	}

	// Register JWT
	token, err := helpers.AuthJWT(infoLogin)
	helpers.PanicError(err)

	return domain.AuthResponse{
		Access_Token: token,
	}
}

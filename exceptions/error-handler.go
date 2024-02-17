package exceptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hasifpriyambudi/cms_test/utils"
)

func ErrorHandler() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		defer func() {
			if reco := recover(); reco != nil {
				switch err := reco.(type) {
				case validator.ValidationErrors:
					res := utils.BuildResponseFailed(http.StatusBadRequest, "", err.Error(), "")
					ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
					return
				case NotFoundError:
					res := utils.BuildResponseFailed(http.StatusNotFound, "", err.Error.Error(), "")
					ctx.AbortWithStatusJSON(http.StatusNotFound, res)
					return
				case ReadJsonError:
					res := utils.BuildResponseFailed(http.StatusBadRequest, "", err.Error.Error(), "")
					ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
					return
				case AuthError:
					res := utils.BuildResponseFailed(http.StatusUnauthorized, "", err.Error.Error(), "")
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
					return
				default:
					defError := reco.(error)
					res := utils.BuildResponseFailed(http.StatusInternalServerError, "", defError.Error(), "")
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
					return
				}
			}
		}()

		ctx.Next()

	}
}

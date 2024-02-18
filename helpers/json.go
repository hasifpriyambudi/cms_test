package helpers

import (
	"github.com/gin-gonic/gin"
)

func ReadJSON(c *gin.Context, res interface{}) error {
	return c.ShouldBindJSON(res)
}

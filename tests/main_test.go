package tests

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	server := gin.Default()
	return server
}

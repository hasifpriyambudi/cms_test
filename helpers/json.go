package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadJSON(c *gin.Context, res interface{}) error {
	return c.ShouldBindJSON(res)
}

func ReturnJSON(w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json")
	// encoder := json.NewEncoder(w)
	// err := encoder.Encode(response)
	// PanicError(err)
}

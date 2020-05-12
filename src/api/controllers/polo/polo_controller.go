package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

//Marco seperti ping pong
func Marco(c *gin.Context) {
	c.JSON(http.StatusOK, polo)
}

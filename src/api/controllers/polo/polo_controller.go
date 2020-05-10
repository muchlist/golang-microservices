package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

//Polo seperti ping pong
func Polo(c *gin.Context) {
	c.JSON(http.StatusOK, polo)
}

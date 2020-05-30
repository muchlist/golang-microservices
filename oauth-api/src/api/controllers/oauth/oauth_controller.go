package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muchlist/golang-microservices/oauth-api/src/api/domain/oauth"
	"github.com/muchlist/golang-microservices/oauth-api/src/api/services"
	"github.com/muchlist/golang-microservices/src/api/utils/errors"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	token, err := services.OauthService.GetAccessToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}

package app

import (
	"github.com/muchlist/golang-microservices/oauth-api/src/api/controllers/oauth"
	"github.com/muchlist/golang-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)

	router.POST("/oauth/access-token", oauth.CreateAccessToken)
	router.GET("/oauth/access-token/:token_id", oauth.GetAccessToken)
}

package app

import (
	"github.com/muchlist/golang-microservices/src/api/controllers/polo"
	"github.com/muchlist/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)
}

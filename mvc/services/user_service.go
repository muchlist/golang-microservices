package services

import (
	"github.com/muchlist/golang-microservices/mvc/domain"
	"github.com/muchlist/golang-microservices/mvc/utils"
)

//GetUser mendapatkan user dari domain
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}

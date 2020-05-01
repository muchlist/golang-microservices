package services

import (
	"github.com/muchlist/golang-microservices/mvc/domain"
	"github.com/muchlist/golang-microservices/mvc/utils"
)

type userService struct{}

var (
	//UserService publik
	UserService userService
)

//GetUser mendapatkan user dari domain
func (u *userService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, err
}

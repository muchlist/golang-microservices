package domain

import (
	"fmt"
	"net/http"

	"github.com/muchlist/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Muchlis", LastName: "Keren", Email: "muchlis.keren@gmail.com"},
	}
	//UserDao public
	UserDao userDaoInterface
)

func init() {
	//untuk keperluan testing sehingga dibuat interface
	UserDao = &userDao{}
}

type userDao struct{}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

//GetUser mendapatkan user dari database berdasarkan userID
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	user := users[userID]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("User %v tidak ditemukan", userID),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}

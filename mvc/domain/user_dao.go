package domain

import (
	"fmt"
	"net/http"

	"github.com/muchlist/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 1, FirstName: "Muchlis", LastName: "Keren", Email: "muchlis.keren@gmail.com"},
	}
)

//GetUser mendapatkan user dari database berdasarkan userID
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	user := users[userID]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v tidak ditemukan", userID),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}

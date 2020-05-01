package services

import (
	"net/http"
	"testing"

	"github.com/muchlist/golang-microservices/mvc/domain"
	"github.com/muchlist/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userID int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	//disini letak mockingnya, pada code asli interface di isi oleh Dao User
	//pada test di isi oleh usersDaoMock
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (u *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "User 0 tidak ditemukan",
		}
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 tidak ditemukan", err.Message)
}

func TestGetUserFoundInDatabase(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		user := &domain.User{
			ID:        userID,
			FirstName: "Muchlis",
			LastName:  "Keren",
			Email:     "muchlis.keren@gmail.com",
		}
		return user, nil
	}

	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Muchlis", user.FirstName)
	assert.EqualValues(t, "Keren", user.LastName)
	assert.EqualValues(t, "muchlis.keren@gmail.com", user.Email)
}

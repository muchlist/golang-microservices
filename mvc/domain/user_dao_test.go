package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	//inisialisasi

	//eksekusi
	user, err := UserDao.GetUser(0)

	//validasi
	assert.Nil(t, user, "we were not expacting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 tidak ditemukan", err.Message)

}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Muchlis", user.FirstName)
	assert.EqualValues(t, "Keren", user.LastName)
	assert.EqualValues(t, "muchlis.keren@gmail.com", user.Email)
}

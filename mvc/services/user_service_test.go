package services

import (
	"net/http"
	"testing"

	"github.com/Djudicael/microserviceMVC/mvc/domain"
	"github.com/Djudicael/microserviceMVC/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	//userMock        userDaoMock
	getUserFunction func(userid int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &userDaoMock{}
}

type userDaoMock struct {
}

func (m *userDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userid int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 was not found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)

}
func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userid int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID: 123,
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)

}

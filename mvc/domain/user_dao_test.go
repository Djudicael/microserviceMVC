package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserfound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	//maintenable
	assert.Nil(t, user, "we are not explecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)

	//complicate to maintain
	if user != nil {
		t.Error("we are not explecting a user with id 0")
	}

	if err == nil {
		t.Error("we were expecting an error when user id is 0")
	}
	if err.StatusCode != http.StatusNotFound {
		t.Error("we were expecting 404 when user is not found")
	}

}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Jud", user.FirstName)
}

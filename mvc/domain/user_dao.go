package domain

import (
	"fmt"
	"net/http"

	"github.com/Djudicael/microserviceMVC/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Jud", LastName: "Dub", Email: "test@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(userID int64) (*User, *utils.ApplicationError)
}
type userDao struct {
}

func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {

	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}

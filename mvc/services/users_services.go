package services

import (
	"github.com/Djudicael/microserviceMVC/mvc/domain"
	"github.com/Djudicael/microserviceMVC/mvc/utils"
)

type userService struct {
}

var (
	UserService userService
)

//GetUser devient une methode du struct UserService
func (u *userService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

package services

import (
	"github.com/Djudicael/microserviceMVC/mvc/domain"
	"github.com/Djudicael/microserviceMVC/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}

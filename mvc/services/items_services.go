package services

import (
	"net/http"

	"github.com/Djudicael/microserviceMVC/mvc/domain"
	"github.com/Djudicael/microserviceMVC/mvc/utils"
)

type itemService struct {
}

var (
	ItemService itemService //permet de le rendre public
)

func (i *itemService) Getitems(ietemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}

}

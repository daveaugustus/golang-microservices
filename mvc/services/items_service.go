package services

import (
	"net/http"

	"github.com/davetweetlive/golang-microservices/mvc/domain"
	"github.com/davetweetlive/golang-microservices/mvc/utils"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {

	return nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}

package services

import (
	"github.com/davetweetlive/golang-microservices/mvc/domain"
	"github.com/davetweetlive/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}

package domain

import (
	"fmt"
	"net/http"

	"github.com/davetweetlive/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{ID: 1, FirstName: "Dave", LastName: "Augustus", Email: "email@email.com"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}

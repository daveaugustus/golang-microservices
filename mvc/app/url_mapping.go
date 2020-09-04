package app

import (
	"github.com/davetweetlive/golang-microservices/mvc/controllers"
)

func mapURLs() {
	router.GET("/users/:user_id", controllers.GetUser)
}

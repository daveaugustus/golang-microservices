package app

import (
	"net/http"

	"github.com/davetweetlive/golang-microservices/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

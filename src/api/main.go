package main

import (
	"fmt"

	"github.com/davetweetlive/golang-microservices/src/api/app"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Port 8080 is already occupied!")
	app.StartApp()
}

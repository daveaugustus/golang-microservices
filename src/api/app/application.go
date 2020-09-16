package app

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapUrls()
	if err := router.Run(":9090"); err != nil {
		panic(err)
	}
}

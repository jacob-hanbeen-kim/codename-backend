package app

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	//define routes
	router.GET("/test", getTests)
	router.POST("/test", postTest)
	//starting server by default localhost :8080
	router.Run()
}

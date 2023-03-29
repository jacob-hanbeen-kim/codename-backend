package app

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	//define routes
	router.GET("/test", getTest)
	router.POST("/test", postTest)
	router.PUT("/test", updateTest)
	//starting server by default localhost :8080
	router.Run()
}

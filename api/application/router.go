package application

import "github.com/gin-gonic/gin"

func router(application *Application, router *gin.Engine) {
	getPing(application, router)
}

func getPing(application *Application, router *gin.Engine) {
	router.GET("/ping", application.ping.Pong)
}

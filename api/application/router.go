package application

import "github.com/gin-gonic/gin"

func router(application *Application, router *gin.Engine) {
	getPing(application, router)
	bookEndPoints(application, router)
}

func getPing(application *Application, router *gin.Engine) {
	router.GET("/ping", application.ping.Pong)
}

func bookEndPoints(application *Application, router *gin.Engine) {
	router.GET("/printer", application.bookPrinter.PrintBook)
}

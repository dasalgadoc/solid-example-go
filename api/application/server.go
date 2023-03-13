package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func BuildEngine(application *Application) *gin.Engine {
	engine := gin.Default()
	router(application, engine)

	fmt.Println("starting application")

	return engine
}

func BuildServer(router http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":8081",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

package application

import "solid-example-go/infrastructure/ping"

type (
	Application struct {
		ping ping.Ping
	}
)

func BuildApplication() *Application {
	return &Application{}
}

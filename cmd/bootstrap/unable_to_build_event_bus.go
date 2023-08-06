package bootstrap

import "fmt"

type UnableToBuildEventBus struct {
	Message string
}

func BuildEvenBusError(eventBusName string) *UnableToBuildEventBus {
	return &UnableToBuildEventBus{
		Message: fmt.Sprintf("unable to build %s", eventBusName),
	}
}

func (e *UnableToBuildEventBus) Error() string {
	return e.Message
}

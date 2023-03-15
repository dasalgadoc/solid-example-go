package events

import "fmt"

type ObserverNoRegistered struct {
	Message string
}

func ObserverUnregistered(observerName string) *ObserverNoRegistered {
	return &ObserverNoRegistered{
		Message: fmt.Sprintf("struct %s is no a registered observer", observerName),
	}
}

func (e *ObserverNoRegistered) Error() string {
	return e.Message
}

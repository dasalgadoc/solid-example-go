package events

type Observer interface {
	Update(event DomainEvent) error
}

var RegisteredObservers = map[string]func() Observer{
	"NewBookLogger": getNewBookLogger,
}

func GetObserver(structName string) (Observer, error) {
	observerStruct, ok := RegisteredObservers[structName]
	if !ok {
		return nil, ObserverUnregistered(structName)
	}
	return observerStruct(), nil
}

func getNewBookLogger() Observer {
	return &NewBookLogger{}
}

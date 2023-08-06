package bootstrap

import (
	"solid-example-go/cmd/configs"
	"solid-example-go/internal/domain/events"
	"solid-example-go/internal/infrastructure/eventbus"
)

var eventBuses = map[string]func(config configs.Config) (events.EventBus, error){
	configs.InMemoryType:   buildInMemoryEventBus,
	configs.ConcurrentType: buildConcurrentEventBus,
}

func GetEventBus(config configs.Config) (events.EventBus, error) {
	eventBus, exists := eventBuses[config.EventBus.Type]
	if !exists {
		return nil, BuildEvenBusError(config.EventBus.Type)
	}

	return eventBus(config)
}

func buildInMemoryEventBus(config configs.Config) (events.EventBus, error) {
	bus, err := eventbus.NewInMemoryEventBus()
	if err != nil {
		return nil, err
	}
	registerObservers(config, bus)
	return bus, nil
}

func buildConcurrentEventBus(config configs.Config) (events.EventBus, error) {
	bus, err := eventbus.NewConcurrentEventBus()
	if err != nil {
		return nil, err
	}
	registerObservers(config, bus)
	return bus, nil
}

func registerObservers(config configs.Config, bus events.EventBus) {
	for _, event := range config.EventBus.ObserversByEvent {
		observers := make([]events.Observer, 0)
		for _, structToAppend := range event.Observers {
			observer, err := events.GetObserver(structToAppend)
			if err == nil {
				observers = append(observers, observer)
			}
		}
		bus.AddObserversToEvent(event.EventName, observers)
	}
}

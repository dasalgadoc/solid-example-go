package eventbus

import "solid-example-go/domain/events"

type InMemoryEventBus struct {
	Observers map[string][]events.Observer
}

func NewInMemoryEventBus() (events.EventBus, error) {
	return &InMemoryEventBus{
		Observers: make(map[string][]events.Observer),
	}, nil
}

func (m *InMemoryEventBus) Notify(event events.DomainEvent) error {
	observers := m.Observers[event.EventName()]
	//if !exists {
	//	return events.EventNotFound(event.EventName())
	//}
	for _, observer := range observers {
		err := observer.Update(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *InMemoryEventBus) AddObserversToEvent(eventName string, observers []events.Observer) {
	m.Observers[eventName] = observers
}

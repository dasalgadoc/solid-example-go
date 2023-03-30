package eventbus

import "solid-example-go/domain/events"

type ConcurrentEventBus struct {
	observers map[string][]events.Observer
	channels  map[string]chan events.DomainEvent
}

func NewConcurrentEventBus() (events.EventBus, error) {
	return &ConcurrentEventBus{
		observers: make(map[string][]events.Observer),
		channels:  make(map[string]chan events.DomainEvent),
	}, nil
}

func (m *ConcurrentEventBus) Notify(event events.DomainEvent) error {
	channel, ok := m.channels[event.EventName()]
	if !ok {
		return events.EventNotFound(event.EventName())
	}
	channel <- event
	return nil
}

func (m *ConcurrentEventBus) AddObserversToEvent(eventName string, observers []events.Observer) {
	m.observers[eventName] = append(m.observers[eventName], observers...)
	m.createChannel(eventName)
}

func (m *ConcurrentEventBus) createChannel(eventName string) {
	_, ok := m.channels[eventName]
	if !ok {
		m.channels[eventName] = make(chan events.DomainEvent)
		go func() {
			for event := range m.channels[eventName] {
				observers, ok := m.observers[event.EventName()]
				if ok {
					for _, observer := range observers {
						observer.Update(event)
					}
				}
			}
		}()
	}
}

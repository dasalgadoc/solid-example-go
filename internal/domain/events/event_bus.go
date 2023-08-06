package events

type EventBus interface {
	Notify(event DomainEvent) error
	AddObserversToEvent(eventName string, observers []Observer)
}

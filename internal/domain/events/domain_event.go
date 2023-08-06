package events

type DomainEvent interface {
	EventName() string
}

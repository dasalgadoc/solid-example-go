package events

type BookCreatedDomainEvent struct {
	ISBN string
}

func NewBookCreatedDomainEvent(ISBN string) BookCreatedDomainEvent {
	return BookCreatedDomainEvent{
		ISBN: ISBN,
	}
}

func (b BookCreatedDomainEvent) EventName() string {
	return "book_created"
}

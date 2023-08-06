package events

import (
	"github.com/stretchr/testify/assert"
	"solid-example-go/internal/domain/book"
	"testing"
	"time"
)

type domainEventTestScenario struct {
	test        *testing.T
	domainEvent DomainEvent
	eventName   string
}

func TestAConcreteDomainEventCanBeUsedThroughItsFather_BookCreated(t *testing.T) {
	s := startDomainEventTestScenario(t)
	s.givenADomainEvent(NewBookCreatedDomainEvent("123A"))
	s.whenEventNameIsRequired()
	s.thenStructImplementsExpectedInterface()
	s.thenStructIsACorrectInstance(BookCreatedDomainEvent{})
	s.thenEventNameIsLikeExpected("book_created")
}

func TestAConcreteDomainEventCanBeUsedThroughItsFather_BookUpdated(t *testing.T) {
	s := startDomainEventTestScenario(t)
	book := book.NewBook("", "", "", "")
	s.givenADomainEvent(NewBookUpdatedDomainEvent(time.Now(), book, book))
	s.whenEventNameIsRequired()
	s.thenStructImplementsExpectedInterface()
	s.thenStructIsACorrectInstance(BookUpdatedDomainEvent{})
	s.thenEventNameIsLikeExpected("book_updated")
}

/*-- steps --*/
func startDomainEventTestScenario(t *testing.T) *domainEventTestScenario {
	t.Parallel()
	return &domainEventTestScenario{
		test: t,
	}
}

func (d *domainEventTestScenario) givenADomainEvent(event DomainEvent) {
	d.domainEvent = event
}

func (d *domainEventTestScenario) whenEventNameIsRequired() {
	d.eventName = d.domainEvent.EventName()
}

func (d *domainEventTestScenario) thenStructImplementsExpectedInterface() {
	assert.Implements(d.test, (*DomainEvent)(nil), d.domainEvent)
}

func (d *domainEventTestScenario) thenStructIsACorrectInstance(expected interface{}) {
	assert.IsType(d.test, expected, d.domainEvent)
}

func (d *domainEventTestScenario) thenEventNameIsLikeExpected(expected string) {
	assert.Equal(d.test, expected, d.eventName)
}

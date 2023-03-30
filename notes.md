# Project notes

## Event bus considerations

I implemented fantastic event bus with goroutines and channels, with some complexity due to coupling with interfaces.

I notice that my interfaces works with some generals queueing systems like RabbitMQ, SQS, In memory implementations, etc.
But, if it is possible, I strongly recommend just implement channels and goroutines for inner event buses. Here an example 
of its simplicity.

### Example 1:
Mono file

```go
package example

type DomainEvent interface {
    EventName() string
}

// A domain event
type UserRegisteredEvent struct {
    UserID string
    Name   string
    Email  string
}

// Channel of this event (you can use interfaces)
var eventBus = make(chan UserRegisteredEvent)
var moreGeneral = make(chan DomainEvent)

func RegisterUser(name string, email string) error {
    userID := generateUserID()
	
	// ellipsis: Code for register user
	
    userRegisteredEvent := UserRegisteredEvent{
        UserID: userID,
        Name:   name,
        Email:  email,
    }
	
    eventBus <- userRegisteredEvent // Domain event to its "observer"
    return nil
}

func SendWelcomeEmail(userID string, email string) error {
    // ellipsis: Observer code...
    return nil
}

func UpdateStats(userID string) error {
    // ellipsis: Observer code...
    return nil
}

func main() {
    // establish a goroutine for observer 
    go func() {
        for {
            event := <-eventBus // dispatch domain event from channel
            SendWelcomeEmail(event.UserID, event.Email)
        }
    }()

    go func() {
        for {
            event := <-eventBus // dispatch domain event from channel
			UpdateStats(event.UserID)
        }
    }()

    // Use case
    RegisterUser("John Doe", "johndoe@example.com")
}
```

### Example 2:
With interfaces

```go
type DomainEvent interface {
    Name() string
}

type EventBus struct {
    channels map[string][]chan DomainEvent
}

func NewEventBus() *EventBus {
    return &EventBus{
        channels: make(map[string][]chan DomainEvent),
    }
}

func (eb *EventBus) Notify(event DomainEvent) error {
    channels, ok := eb.channels[event.Name()]
    if !ok {
        return fmt.Errorf("no observers registered for event %s", event.Name())
    }

    for _, ch := range channels {
        go func(ch chan DomainEvent) {
            ch <- event
        }(ch)
    }
    return nil
}

func (eb *EventBus) AddObserverToEvent(eventName string, observer chan DomainEvent) {
    eb.channels[eventName] = append(eb.channels[eventName], observer)
}

// Use case
func sendWelcomeEmail(event DomainEvent) error {
    newUser, ok := event.(NewUserEvent)
    if !ok {
        return fmt.Errorf("invalid event type")
    }

    // No code - just a concept
    err := concreteWelcomeEmail(newUser.Email)
    if err != nil {
        return err
    }
    return nil
}

func main() {
	// Builder section
    bus := NewEventBus()
    bus.AddObserverToEvent("NewUserEvent", sendWelcomeEmail)

    newUser := &User{Nombre: "Juan", Apellido: "Pérez", Email: "juan.perez@example.com"}
	
    err := concreteUserRegister(newUser)
    if err != nil {
        log.Fatal(err)
    }
}
```
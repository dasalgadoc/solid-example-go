package notification

import (
	"fmt"
	"solid-example-go/domain/events"
	"solid-example-go/domain/notification"
)

const MESSAGE_FORM = "New book (%s) created!"

type BookCreatedNotificationHandler struct {
	NotificationSenders []notification.MultipleSender
}

func NewSlackNotificationHandler(
	senders []notification.MultipleSender,
) BookCreatedNotificationHandler {
	return BookCreatedNotificationHandler{
		NotificationSenders: senders,
	}
}

func (s *BookCreatedNotificationHandler) Update(event events.DomainEvent) error {
	bookCreatedDomainEvent, ok := event.(events.BookCreatedDomainEvent)
	if !ok {
		return events.EventCastFail(event.EventName(), "book_created")
	}
	subject, err := notification.NewNotificationSubject(bookCreatedDomainEvent.EventName())
	if err != nil {
		return err
	}
	content := notification.NewNotificationContent(fmt.Sprintf(MESSAGE_FORM, bookCreatedDomainEvent.ISBN))

	for _, sender := range s.NotificationSenders {
		sender.SendToMultipleDestination(subject, content)
	}

	return nil
}

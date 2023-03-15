package notification

import (
	"fmt"
	"solid-example-go/domain/events"
	"solid-example-go/domain/notification"
)

const MESSAGE_FORM = "New book (%s) created!"

type SlackNotificationHandler struct {
	SlackNotificationSender SlackNotificationSender
	SlackChannels           []notification.ReceiverSlackChannel
}

func NewSlackNotificationHandler(
	sender SlackNotificationSender,
	channels []notification.ReceiverSlackChannel) SlackNotificationHandler {
	return SlackNotificationHandler{
		SlackNotificationSender: sender,
		SlackChannels:           channels,
	}
}

func (s *SlackNotificationHandler) Update(event events.DomainEvent) error {
	bookCreatedDomainEvent, ok := event.(events.BookCreatedDomainEvent)
	if !ok {
		return events.EventCastFail(event.EventName(), "book_created")
	}
	subject, err := notification.NewNotificationSubject(bookCreatedDomainEvent.EventName())
	if err != nil {
		return err
	}
	content := notification.NewNotificationContent(fmt.Sprintf(MESSAGE_FORM, bookCreatedDomainEvent.ISBN))

	for _, slackChannel := range s.SlackChannels {
		s.SlackNotificationSender.SendNotification(&slackChannel, subject, content)
	}
	return nil
}

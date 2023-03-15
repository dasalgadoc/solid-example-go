package notification

import "solid-example-go/domain/notification"

type SlackNotificationSender struct {
	SlackChannels []notification.ReceiverSlackChannel
}

func NewSlackNotificationSender(channels []notification.ReceiverSlackChannel) SlackNotificationSender {
	return SlackNotificationSender{
		SlackChannels: channels,
	}
}

func (s *SlackNotificationSender) SendNotification(
	receiver notification.NotificationReceiver,
	subject notification.NotificationSubject,
	content notification.NotificationContect) {
	// TODO
}

func (s *SlackNotificationSender) SendToMultipleDestination(
	subject notification.NotificationSubject, content notification.NotificationContect) {
	for _, slackChannel := range s.SlackChannels {
		s.SendNotification(&slackChannel, subject, content)
	}
}

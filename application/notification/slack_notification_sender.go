package notification

import "solid-example-go/domain/notification"

type SlackNotificationSender struct{}

func (s *SlackNotificationSender) SendNotification(
	receiver notification.NotificationReceiver,
	subject notification.NotificationSubject,
	content notification.NotificationContect) {
	// TODO
}

package notification

type Notificable interface {
	SendNotification(receiver NotificationReceiver, subject NotificationSubject, content NotificationContect)
}

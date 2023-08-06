package notification

type MultipleSender interface {
	SendToMultipleDestination(subject NotificationSubject, content NotificationContect)
}

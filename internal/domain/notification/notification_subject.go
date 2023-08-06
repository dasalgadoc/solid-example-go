package notification

type NotificationSubject struct {
	Subject string
}

func NewNotificationSubject(subject string) (NotificationSubject, error) {
	if len(subject) > 140 {
		return NotificationSubject{}, SubjectTooLong()
	}
	return NotificationSubject{
		Subject: subject,
	}, nil
}

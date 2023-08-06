package notification

type NotificationContect struct {
	Content string
}

func NewNotificationContent(content string) NotificationContect {
	return NotificationContect{
		Content: content,
	}
}

package notification

type NotificationReceiver interface {
	Receiver() string
	IsReceiverValid() bool
}

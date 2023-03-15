package notification

type ReceiverNotValid struct {
	Message string
}

func ReceiverIsNotValid() *ReceiverNotValid {
	return &ReceiverNotValid{
		Message: "receiver is not valid",
	}
}

func (e *ReceiverNotValid) Error() string {
	return e.Message
}

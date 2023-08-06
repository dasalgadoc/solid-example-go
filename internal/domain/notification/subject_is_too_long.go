package notification

type SubjectIsTooLong struct {
	Message string
}

func SubjectTooLong() *SubjectIsTooLong {
	return &SubjectIsTooLong{
		Message: "subject is too long, it has to be 140 characters or less",
	}
}

func (e *SubjectIsTooLong) Error() string {
	return e.Message
}

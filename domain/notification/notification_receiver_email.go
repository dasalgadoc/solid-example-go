package notification

import "regexp"

const EMAIL_PATTERN = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

type ReceiverEmail struct {
	Email string
}

func NewReceiverEmail(email string) (ReceiverEmail, error) {
	r := ReceiverEmail{Email: email}
	if !r.IsReceiverValid() {
		return ReceiverEmail{}, ReceiverIsNotValid()
	}
	return r, nil
}

func (e *ReceiverEmail) Receiver() string {
	return e.Email
}

func (e *ReceiverEmail) IsReceiverValid() bool {
	match, _ := regexp.MatchString(EMAIL_PATTERN, e.Email)
	return match
}

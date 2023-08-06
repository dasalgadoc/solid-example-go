package notification

type ReceiverSlackChannel struct {
	SlackChannel string
}

func NewReceiverSlackChannel(slackChannel string) (ReceiverSlackChannel, error) {
	r := ReceiverSlackChannel{SlackChannel: slackChannel}
	if !r.IsReceiverValid() {
		return ReceiverSlackChannel{}, ReceiverIsNotValid()
	}
	return r, nil
}

func (s *ReceiverSlackChannel) Receiver() string {
	return s.SlackChannel
}

func (s *ReceiverSlackChannel) IsReceiverValid() bool {
	return s.SlackChannel != ""
}

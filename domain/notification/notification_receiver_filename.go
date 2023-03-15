package notification

type ReceiverFilename struct {
	Path string
}

func NewReceiverFilename(path string) (ReceiverFilename, error) {
	r := ReceiverFilename{Path: path}
	if !r.IsReceiverValid() {
		return ReceiverFilename{}, ReceiverIsNotValid()
	}
	return r, nil
}

func (f *ReceiverFilename) Receiver() string {
	return f.Path
}

func (f *ReceiverFilename) IsReceiverValid() bool {
	return f.Path != ""
}

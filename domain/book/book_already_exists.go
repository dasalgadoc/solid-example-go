package book

type BookAlreadyExists struct {
	Message string
}

func AlreadyExists() *BookAlreadyExists {
	return &BookAlreadyExists{
		Message: "book already exists in repository",
	}
}

func (e *BookAlreadyExists) Error() string {
	return e.Message
}

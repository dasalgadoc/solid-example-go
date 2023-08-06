package book

type BookDoesNotExist struct {
	Message string
}

func DoesNotExist() *BookDoesNotExist {
	return &BookDoesNotExist{
		Message: "book does not exist in repository",
	}
}

func (e *BookDoesNotExist) Error() string {
	return e.Message
}

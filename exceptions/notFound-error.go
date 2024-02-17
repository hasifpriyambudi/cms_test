package exceptions

type NotFoundError struct {
	Error error
}

func NewNotFoundError(error error) NotFoundError {
	return NotFoundError{Error: error}
}

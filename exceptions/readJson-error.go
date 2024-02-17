package exceptions

type ReadJsonError struct {
	Error error
}

func NewReadJsonError(error error) ReadJsonError {
	return ReadJsonError{Error: error}
}

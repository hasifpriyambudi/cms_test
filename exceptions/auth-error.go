package exceptions

type AuthError struct {
	Error error
}

func NewAuthError(error error) AuthError {
	return AuthError{Error: error}
}

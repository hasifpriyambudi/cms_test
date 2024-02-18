package exceptions

type MysqlError struct {
	Error error
}

func NewMysqlError(error error) MysqlError {
	return MysqlError{Error: error}
}

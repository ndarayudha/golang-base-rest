package exception

type AlredyExistError struct {
	Error string
}

func NewAlreadyExistError(error string) AlredyExistError {
	return AlredyExistError{Error: error}
}

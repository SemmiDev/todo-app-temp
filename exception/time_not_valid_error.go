package exception

type TimeNotValidError struct {
	Error string
}

func NewTimeNotValidError(error string) TimeNotValidError {
	return TimeNotValidError{Error: error}
}

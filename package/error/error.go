package error

type CError struct {
	Message string
	Code    int
}

func (e CError) Error() string {
	return e.Message
}

func NewCError(code int, message string) CError {
	return CError{
		Code:    code,
		Message: message,
	}
}

package exceptions

type ValidationError struct {
	Message string
	Code    int
}

func NewValidationError(message string) ValidationError {
	return ValidationError{
		Message: message,
		Code:    422,
	}
}

func (v ValidationError) Error() string {
	return v.Message
}

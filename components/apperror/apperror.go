package apperror

import "fmt"

type AppError struct {
	StatusCode int    `json:"status"`
	ErrorType  string `json:"error_type"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (ae *AppError) Error() string {
	if ae.Err != nil {
		return fmt.Sprintf("%s: %s | internal: %v", ae.ErrorType, ae.Message, ae.Err)
	}
	return fmt.Sprintf("%s: %s", ae.ErrorType, ae.Message)
}

func NewAppError(statusCode int, errorType, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		ErrorType:  errorType,
		Message:    message,
		Err:        err,
	}
}

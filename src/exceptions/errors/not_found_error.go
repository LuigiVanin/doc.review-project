package errors

type NotFoundError = AppError

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Status:  404,
		Message: message,
		Code:    "NOT_FOUND_ERROR",
	}
}

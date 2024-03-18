package errors

type InternalServerError = AppError

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{
		Status:  500,
		Message: message,
		Code:    "INTERNAL_SERVER_ERROR",
	}
}

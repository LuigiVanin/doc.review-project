package errors

type ForbiddenError = AppError

func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{
		Status:  403,
		Message: message,
		Code:    "INTERNAL_SERVER_ERROR",
	}
}

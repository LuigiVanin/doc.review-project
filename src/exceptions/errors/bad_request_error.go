package errors

type BadRequestError = AppError

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Status:  400,
		Message: message,
		Code:    "BAD_REQUEST_ERROR",
	}
}

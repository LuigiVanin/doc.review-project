package errors

type UnauthorizedError = AppError

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Status:  401,
		Message: message,
		Code:    "UNAUTHORIZED_ERROR",
	}
}

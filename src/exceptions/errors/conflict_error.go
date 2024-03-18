package errors

type ConflictError = AppError

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Status:  409,
		Message: message,
		Code:    "CONFLICT_ERROR",
	}
}

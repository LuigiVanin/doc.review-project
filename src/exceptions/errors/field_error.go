package errors

type FieldError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

type FieldErrors struct {
	Fields []FieldError
}

func (e *FieldErrors) Error() string {
	errorText := ""
	for _, error := range e.Fields {
		errorText += error.Field + " " + error.Tag + "\n"
	}
	return errorText
}

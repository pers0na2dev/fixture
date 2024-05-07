package fixture

type FieldError struct {
	FieldName string
	Message   string
}

func (e *FieldError) Error() string {
	return e.FieldName + ": " + e.Message
}
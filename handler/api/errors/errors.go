package errors

type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func New(text string) error {
	return &Error{
		Message: text,
	}
}

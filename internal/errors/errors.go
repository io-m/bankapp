package errors

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewError(m string, code int) *Error {
	return &Error{
		Message:    m,
		StatusCode: code,
	}
}
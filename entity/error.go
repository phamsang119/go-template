package entity

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewFromError(code int, err error) *Error {
	return &Error{
		Code:    code,
		Message: err.Error(),
	}
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

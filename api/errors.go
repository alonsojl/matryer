package api

import "net/http"

type Error struct {
	Code int
	Err  string
}

func (e Error) Error() string {
	return e.Err
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "invalid request'",
	}
}

func ErrInternalServer() Error {
	return Error{
		Code: http.StatusInternalServerError,
		Err:  "internal server",
	}
}

func ErrUnauthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err:  "unauthorized request",
	}
}

package HttpError

import (
	"fmt"
)

// MARK: Internal Server Error
type InternalServerError struct {
	AdditionalData interface{}
}

func (e *InternalServerError) Error() string {
	return "internal server error"
}

func (e *InternalServerError) RestCode() int {
	return 500
}

// MARK: Page Not Found Error
type PageNotFoundError struct {
	Method string
	Path   string
}

func (e *PageNotFoundError) Error() string {
	arg := ""
	if e.Path != "" {
		arg = fmt.Sprintf("%v %v: ", e.Method, e.Path)
	}
	return arg + " page not found"
}

func (e *PageNotFoundError) RestCode() int {
	return 404
}

// MARK: Bad Request Error
type BadRequestError struct {
	Msg string
}

func (e *BadRequestError) Error() string {
	return "bad request: " + e.Msg
}

func (e *BadRequestError) RestCode() int {
	return 400
}

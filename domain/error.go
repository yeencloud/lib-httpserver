package domain

import (
	"fmt"
)

type RestErrorCode interface {
	RestCode() int
}

type InternalErrorCode interface {
	InternalErrorCode() string
}

type InternalServerError struct {
	AdditionalData interface{}
}

func (e *InternalServerError) Error() string {
	return "http: internal server error"
}

func (e *InternalServerError) RestCode() int {
	return 500
}

type ResponseError struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message"`
	Translation string `json:"translation,omitempty"`
}

type PageNotFoundError struct {
	Msg string

	Method string
	Path   string
}

func (e *PageNotFoundError) Error() string {
	arg := ""
	if e.Path != "" {
		arg = fmt.Sprintf("(%v %v)", e.Method, e.Path)
	}
	return fmt.Sprintf("http: page not found: %v %v", e.Msg, arg)
}

func (e *PageNotFoundError) RestCode() int {
	return 404
}

type BadRequestError struct {
	Msg string
}

func (e *BadRequestError) Error() string {
	return "http: bad request: " + e.Msg
}

func (e *BadRequestError) RestCode() int {
	return 400
}

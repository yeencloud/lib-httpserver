package domain

import (
	"fmt"

	"github.com/yeencloud/lib-shared/apperr"
)

type PageNotFoundError struct {
	Method string
	Path   string
}

func (p PageNotFoundError) Error() string {
	return fmt.Sprintf("%s %s not found", p.Method, p.Path)
}

func (p PageNotFoundError) Unwrap() error {
	return apperr.ResourceNotFoundError{}
}

func NewPageNotFoundError(method string, path string) PageNotFoundError {
	return PageNotFoundError{
		Method: method,
		Path:   path,
	}
}

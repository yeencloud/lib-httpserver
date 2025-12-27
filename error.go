package httpserver

import (
	"net/http"

	"github.com/yeencloud/lib-shared/apperr"
)

func mapErrorTypeToHttpCode(err error) int {
	errtype := apperr.GetErrorTypeOrNil(err)
	const defaultError = http.StatusInternalServerError

	if errtype == nil {
		return defaultError
	}

	switch *errtype {
	case apperr.ErrorTypeInvalidArgument:
		return http.StatusBadRequest
	case apperr.ErrorTypeResourceNotFound:
		return http.StatusNotFound
	case apperr.ErrorTypeNotImplemented:
		return http.StatusNotImplemented
	case apperr.ErrorTypeUnavailable:
		return http.StatusServiceUnavailable
	case apperr.ErrorTypeUnauthorized:
		return http.StatusUnauthorized
	default:
		return defaultError
	}
}

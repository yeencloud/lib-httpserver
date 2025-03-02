package httpserver

import (
	"errors"
	"os"
	"strings"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"

	"github.com/yeencloud/httpserver/domain"
)

type Response struct {
	StatusCode int `json:"status"`

	Body  interface{}           `json:"body,omitempty"`
	Error *domain.ResponseError `json:"error,omitempty"`

	RequestID string `json:"requestId,omitempty"`
}

func reply(ctx *gin.Context, replyCall func(code int, obj any), code int, body interface{}, err error) {
	if ctx.Writer.Written() {
		return
	}

	var response any

	if raw, _ := ctx.Get("raw_response"); raw != nil && raw.(bool) == true {
		if err != nil {
			body = err.Error()
		}

		response = body
	} else {
		response = Response{
			StatusCode: code,
			Body:       body,
		}

		if structs.IsZero(body) {
			response.(*Response).Body = nil
		}

		if err != nil {
			errorStr := err.Error()
			errs := strings.Split(errorStr, "\n")

			if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "prod" {
				if len(errs) > 1 {
					errorStr = errs[0]
				}
			}

			response.(*Response).Error = &domain.ResponseError{
				Message: errorStr,
			}
		}
	}

	replyCall(code, response)
}

type renderFunc func(code int, obj any)

func (hs *HttpServer) renderFunc(ctx *gin.Context) renderFunc {
	value, ok := ctx.Get("responseType")
	if ok {
		if value == "xml" {
			return ctx.XML
		}
	}

	return ctx.JSON
}

func (hs *HttpServer) Reply(ctx *gin.Context, body interface{}) {
	reply(ctx, hs.renderFunc(ctx), 200, body, nil)
}

func (hs *HttpServer) ReplyWithError(ctx *gin.Context, err error, body ...interface{}) {
	code := 500
	var restError domain.RestErrorCode
	if err != nil && errors.As(err, &restError) {
		code = restError.RestCode()
	}
	reply(ctx, hs.renderFunc(ctx), code, body, err)
}

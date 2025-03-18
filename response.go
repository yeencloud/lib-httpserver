package httpserver

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yeencloud/lib-httpserver/domain"
)

type Response struct {
	StatusCode int `json:"status"`

	Body  interface{}           `json:"body,omitempty"`
	Error *domain.ResponseError `json:"error,omitempty"`

	RequestId string `json:"requestId,omitempty"`
	ContextId string `json:"contextId,omitempty"`
}

func reply(ctx *gin.Context, replyCall func(code int, obj any), code int, body interface{}, err error) {
	if ctx.Writer.Written() {
		return
	}

	response := Response{
		StatusCode: code,
		Body:       body,
		ContextId:  ctx.GetString("context_id"),
		RequestId:  ctx.GetString("request_id"),
	}

	if err != nil {
		errorStr := err.Error()
		errs := strings.Split(errorStr, "\n")

		if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "prod" { // TODO check env differently
			if len(errs) > 1 {
				errorStr = errs[0]
			}
		}

		response.Error = &domain.ResponseError{
			Message: errorStr,
		}
	}

	replyCall(code, response)
}

type renderFunc func(code int, obj any)

func (hs *HttpServer) renderFunc(ctx *gin.Context) renderFunc {
	if ctx.Writer.Header().Get("Content-Type") == gin.MIMEXML {
		return ctx.XML
	}
	return ctx.JSON
}

func (hs *HttpServer) Reply(ctx *gin.Context, body interface{}) {
	code := http.StatusOK
	if ctx.Request.Method == http.MethodPost {
		code = http.StatusCreated
	}
	reply(ctx, hs.renderFunc(ctx), code, body, nil)
}

func (hs *HttpServer) ReplyWithError(ctx *gin.Context, err error) {
	code := 500
	var restError domain.RestErrorCode
	if err != nil && errors.As(err, &restError) {
		code = restError.RestCode()
	}

	reply(ctx, hs.renderFunc(ctx), code, nil, err)
}

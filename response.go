package httpserver

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yeencloud/lib-httpserver/domain"
	"github.com/yeencloud/lib-httpserver/domain/error"
	sharedErrors "github.com/yeencloud/lib-shared/errors"
)

func (hs *HttpServer) reply(ctx *gin.Context, replyCall func(code int, obj any), code int, body interface{}, err error) {
	if ctx.Writer.Written() {
		return
	}

	response := domain.Response{
		StatusCode:    code,
		CorrelationId: GetCorrelationID(ctx),
		RequestId:     GetRequestID(ctx),
	}

	if err != nil {
		errorStr := err.Error()
		errs := strings.Split(errorStr, "\n")

		if hs.Env.IsProduction() && len(errs) > 1 {
			errorStr = errs[0]
		}

		response.Error = &domain.ResponseError{
			Message: errorStr,
		}

		var IdentifierError sharedErrors.IdentifiableError
		var FixError sharedErrors.FixableError

		if errors.As(err, &IdentifierError) {
			response.Error.Code = IdentifierError.Identifier()
		}

		if errors.As(err, &FixError) {
			response.Error.HowToFix = FixError.HowToFix()
		}
	} else {
		response.Body = body
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
	hs.reply(ctx, hs.renderFunc(ctx), code, body, nil)
}

func (hs *HttpServer) ReplyWithError(ctx *gin.Context, err error) {
	code := http.StatusInternalServerError

	var restError HttpError.RestErrorCode
	if err != nil && errors.As(err, &restError) {
		code = restError.RestCode()
	}

	hs.reply(ctx, hs.renderFunc(ctx), code, nil, err)
}

package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/samber/lo"

	"github.com/gin-gonic/gin"
	"github.com/yeencloud/lib-httpserver/contract/httpserver"

	"github.com/yeencloud/lib-httpserver/domain/error"
	sharedErrors "github.com/yeencloud/lib-shared/errors"
)

func structToMap(v any) (map[string]interface{}, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (hs *HttpServer) reply(ctx *gin.Context, replyCall func(code int, obj any), code int, body any, err error) {
	if ctx.Writer.Written() {
		return
	}

	response := httpserver.Response{
		Status:        code,
		CorrelationId: lo.ToPtr(GetCorrelationID(ctx)),
		RequestId:     lo.ToPtr(GetRequestID(ctx)),
	}

	if err != nil {
		hs.fillResponseWithErrorDetails(err, &response)
	} else {
		structmap, err := structToMap(body)
		if err != nil {
			hs.reply(ctx, replyCall, http.StatusInternalServerError, nil, err)
			return
		} else {
			response.Body = lo.ToPtr(structmap)
		}
	}

	replyCall(code, response)
}

func (hs *HttpServer) fillResponseWithErrorDetails(err error, response *httpserver.Response) {
	errorStr := err.Error()
	errs := strings.Split(errorStr, "\n")

	if hs.Env.IsProduction() && len(errs) > 1 {
		errorStr = errs[0]
	}

	response.Error = &httpserver.ResponseError{
		Message: errorStr,
	}

	var IdentifierError sharedErrors.IdentifiableError
	var FixError sharedErrors.FixableError

	if errors.As(err, &IdentifierError) {
		response.Error.Code = lo.ToPtr(IdentifierError.Identifier())
	}

	if errors.As(err, &FixError) {
		response.Error.HowToFix = lo.ToPtr(FixError.HowToFix())
	}
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

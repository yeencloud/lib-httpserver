package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/samber/lo"
	"github.com/yeencloud/lib-httpserver/domain"

	"github.com/gin-gonic/gin"

	appErrors "github.com/yeencloud/lib-shared/apperr"
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

	response := domain.Response{
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

func (hs *HttpServer) fillResponseWithErrorDetails(err error, response *domain.Response) {
	errorStr := err.Error()
	errs := strings.Split(errorStr, "\n")

	if hs.Env.IsProduction() && len(errs) > 1 {
		errorStr = errs[0]
	}

	response.Error = &domain.ResponseError{
		Message: errorStr,
	}

	var FixError appErrors.FixableError

	if errors.As(err, &FixError) {
		response.Error.TroubleshootingTip = lo.ToPtr(FixError.TroubleshootingTip())
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
	code := mapErrorTypeToHttpCode(err)
	hs.reply(ctx, hs.renderFunc(ctx), code, nil, err)
}

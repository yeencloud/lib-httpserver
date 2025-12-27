package domain

type Response struct {
	Body          *map[string]interface{} `json:"body,omitempty"`
	CorrelationId *string                 `json:"correlationId,omitempty"`
	Error         *ResponseError          `json:"error,omitempty"`
	RequestId     *string                 `json:"requestId,omitempty"`
	Status        int                     `json:"status"`
}
type ResponseError struct {
	Code               *string `json:"code"`
	TroubleshootingTip *string `json:"troubleshooting"`
	Message            string  `json:"message"`
}

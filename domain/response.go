package domain

type ResponseError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`

	HowToFix string `json:"howToFix,omitempty"`
}

type Response struct {
	StatusCode int `json:"status"`

	Body  interface{}    `json:"body,omitempty"`
	Error *ResponseError `json:"error,omitempty"`

	RequestId     string `json:"requestId,omitempty"`
	CorrelationId string `json:"correlationId,omitempty"`
}

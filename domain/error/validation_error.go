package HttpError

// MARK: CorrelationID Required Error
type CorrelationIDRequiredError struct {
}

func (e *CorrelationIDRequiredError) Error() string {
	return "Correlation ID is required"
}

func (e *CorrelationIDRequiredError) Unwrap() error {
	return &BadRequestError{}
}

func (e *CorrelationIDRequiredError) HowToFix() string {
	return "Specify the header 'X-Correlation-ID' in the request as a valid UUID"
}

func (e *CorrelationIDRequiredError) Identifier() string {
	return "HTTP-REQ-001"
}

// MARK: RequestID Required Error
type RequestIDRequiredError struct {
}

func (e *RequestIDRequiredError) Error() string {
	return "Request ID is required"
}

func (e *RequestIDRequiredError) Unwrap() error {
	return &BadRequestError{}
}

func (e *RequestIDRequiredError) HowToFix() string {
	return "Specify the header 'X-Request-ID' in the request as a valid UUID"
}

func (e *RequestIDRequiredError) Identifier() string {
	return "HTTP-REQ-002"
}

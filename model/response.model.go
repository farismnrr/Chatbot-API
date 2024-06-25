package model

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{Code: code, Message: message}
}

func NewSuccessResponse(code int, message string) *SuccessResponse {
	return &SuccessResponse{Code: code, Message: message}
}

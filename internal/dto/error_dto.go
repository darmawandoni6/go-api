package dto

// ErrorResponse is the standard envelope for all error responses.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SuccessResponse is the standard envelope for all success responses.
// Data holds the actual payload (any struct, slice, etc.).
type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

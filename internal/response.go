package errors

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Extra   interface{} `json:"details,omitempty"`
}

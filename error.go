package notify

// APIError is a custom error-type struct adjusted for the API usage.
type APIError struct {
	Message    string
	StatusCode int
	Errors     []Error
}

// Error method is here to return a top level error message as well as define
// our new error type.
func (e *APIError) Error() string {
	return e.Message
}

// Error may be returned by the API.
type Error struct {
	Error   string
	Message string
}

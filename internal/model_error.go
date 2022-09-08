package internal

// Internal error object model.
type ModelError struct {

	// A short error code that defines the error, meant for programmatic parsing.
	Code string `json:"code"`

	// A human-readable error string.
	Message string `json:"message"`
}

package response

// Error error
type Error struct {
	Code    int    `json:"code,omitempty" mapstructure:"code"`
	Message string `json:"message,omitempty" mapstructure:"message"`
}

// Results return results
type Results struct {
}

// Error error
func (ec Error) Error() string {
	return ec.Message
}

// ErrorCode get error code
func (ec Error) ErrorCode() int {
	return ec.Code
}

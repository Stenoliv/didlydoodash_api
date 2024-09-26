package utils

// EndpointError defines the structure of an endpoint error response.
type EndpointError struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Tag    string `json:"message"`
}

// NewEndpointError creates a new EndpointError with the given code and message.
func NewEndpointError(code int, tag string) *EndpointError {
	return &EndpointError{Status: "error", Code: code, Tag: tag}
}

// Error codes
const (
	// User input error codes
	CodeInvalidInput    = 100
	CodeInvalidUsername = 101
	CodeInvalidPassword = 102

	// Authentication error codes
	CodeUserNotFound        = 200
	CodeAuthenticationError = 201
	CodeTokenError          = 202
	CodeInvalidToken        = 203
	CodeNotAuthenticated    = 204
	CodeNotVerified         = 205

	// Too many requests
	CodeTooManyRequests = 400

	// Server error codes
	CodeServerError   = 500
	CodeSaveUserError = 501
)

// User input errors
var (
	InvalidInput    = NewEndpointError(CodeInvalidInput, "Invalid input")
	InvalidUsername = NewEndpointError(CodeInvalidUsername, "Invalid Username")
	InvalidPassword = NewEndpointError(CodeInvalidPassword, "Wrong Password")
)

// Authentication errors
var (
	UserNotFound        = NewEndpointError(CodeUserNotFound, "Account not found")
	AuthenticationError = NewEndpointError(CodeAuthenticationError, "Failed to authenticate")
	TokenError          = NewEndpointError(CodeTokenError, "Failed to generate token")
	InvalidToken        = NewEndpointError(CodeInvalidToken, "JWT not valid")
	NotAuthenticated    = NewEndpointError(CodeNotAuthenticated, "You are not authenticated")
	NotVerified         = NewEndpointError(CodeNotVerified, "You are not verified")
)

// Server errors
var (
	ServerError   = NewEndpointError(CodeServerError, "Error occurred on server")
	SaveUserError = NewEndpointError(CodeSaveUserError, "Error on registering user")
)

// Too many requests errors
var (
	TooManyRequests = NewEndpointError(CodeTooManyRequests, "Too many requests")
)

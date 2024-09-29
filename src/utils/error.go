package utils

import (
	"log"
	"time"
)

// LogError logs errors with a timestamp to the console
func LogError(err error, message string) {
	if err != nil {
		log.Printf("[%s] ERROR: %s - %v\n", time.Now().Format(time.RFC3339), message, err)
	}
}

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
	CodeInvalidInput = 100

	// Authentication error codes
	CodeUserNotFound        = 200
	CodeAuthenticationError = 201
	CodeTokenError          = 202
	CodeInvalidToken        = 203
	CodeNotAuthenticated    = 204
	CodeNotVerified         = 205

	// Organisation errors: 300
	CodeOrgCreate = 300

	CodeChatNotFound = 350

	// Server error codes
	CodeServerError   = 500
	CodeSaveUserError = 501
)

// User input errors
var (
	InvalidInput = NewEndpointError(CodeInvalidInput, "Invalid input")
)

// Authentication errors
var (
	UserNotFound        = NewEndpointError(CodeUserNotFound, "Account not found")
	AuthenticationError = NewEndpointError(CodeAuthenticationError, "Failed to authenticate")
	InvalidToken        = NewEndpointError(CodeInvalidToken, "JWT not valid")
	NotAuthenticated    = NewEndpointError(CodeNotAuthenticated, "You are not authenticated")
)

// Organisation errors
var (
	FailedToCreateOrg = NewEndpointError(CodeOrgCreate, "Failed to create organisation try again later")

	ChatNotFound = NewEndpointError(CodeChatNotFound, "Chat not found")
)

// Server errors
var (
	ServerError   = NewEndpointError(CodeServerError, "Error occurred on server")
	SaveUserError = NewEndpointError(CodeSaveUserError, "Error on registering user")
)

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
	CodeAuthorityError      = 203
	CodeTokenError          = 202
	CodeInvalidToken        = 203
	CodeNotAuthenticated    = 204
	CodeNotVerified         = 205

	// Organisation errors: 300
	CodeOrgCreate      = 300
	CodeOrgNotFound    = 304
	CodeMemberNotFound = 310

	// Organisation chats
	CodeChatNotFound       = 330
	CodeChatMemberNotFound = 334

	// Project error: 350
	CodeProjectCreate         = 350
	CodeProjectNotFound       = 354
	CodeProjectMemberNotFound = 364

	// Kanban error: 380
	CodeKanbanCreate   = 380
	CodeKanbanNotFound = 384

	// Whiteboard error: 700
	CodeWhiteboardCreate = 700
	CodeWhiteboardNotFound = 704


	// Server error codes
	CodeServerError   = 500
	CodeSaveUserError = 501

	// Websocket errors
	CodeWebsocketFailed = 600
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
	NotEnoughAuthority  = NewEndpointError(CodeAuthorityError, "You are forbidden to use this")
)

// Organisation errors
var (
	FailedToCreateOrg = NewEndpointError(CodeOrgCreate, "Failed to create organisation try again later")
	OrgNotFound       = NewEndpointError(CodeOrgNotFound, "Organisation does not exist")
	MemberNotFound    = NewEndpointError(CodeMemberNotFound, "Member was not found in organisation")

	// Chat errors
	ChatNotFound       = NewEndpointError(CodeChatNotFound, "Chat not found")
	ChatMemberNotFound = NewEndpointError(CodeChatMemberNotFound, "Not part of chat")

	// Project errors
	ProjectCreateError    = NewEndpointError(CodeProjectCreate, "Failed to create project")
	ProjectNotFound       = NewEndpointError(CodeProjectNotFound, "Project not found")
	ProjectMemberNotFound = NewEndpointError(CodeProjectMemberNotFound, "Project member not found")
)

// Kanban errors
var (
	KanbanCreateError = NewEndpointError(CodeKanbanCreate, "Failed to create kanban")
	KanbanNotFound    = NewEndpointError(CodeKanbanNotFound, "Kanban not found")
)

// Whiteboard errors 
var (
	WhiteboardCreateError = NewEndpointError(CodeWhiteboardCreate, "Failed to create Whiteboard")
	WhiteboardNotFound    = NewEndpointError(CodeWhiteboardNotFound, "Whiteboard not found")
)

// WebSocket errors
var (
	WebSocketFailed = NewEndpointError(CodeWebsocketFailed, "Failed to connect to websocket")
)

// Server errors
var (
	ServerError   = NewEndpointError(CodeServerError, "Error occurred on server")
	SaveUserError = NewEndpointError(CodeSaveUserError, "Error on registering user")
)

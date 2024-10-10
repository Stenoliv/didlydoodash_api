package chat

import (
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	Hub             *Hub
	NotificationHub *NotificationHub
}

func NewChatHandler() *ChatHandler {
	hub := NewHub()
	notificationHub := NewChatHubSSE()
	return &ChatHandler{Hub: hub, NotificationHub: notificationHub}
}

// Join chat room
func (h *ChatHandler) JoinRoom(c *gin.Context) {
	roomID := c.Param("chatId")
	// Retrive chat in database and check that user is part of
	room, err := daos.GetChat(roomID)
	if err != nil || room == nil {
		c.JSON(http.StatusBadRequest, utils.ChatNotFound)
		return
	}

	ok := false
	for _, member := range room.Members {
		if member.UserID == *models.CurrentUser {
			ok = true
		}
	}
	if !ok {
		c.JSON(http.StatusForbidden, utils.ChatMemberNotFound)
		return
	}

	// Create websocket connection
	conn, err := ws.WebsocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.WebSocketFailed)
		return
	}

	// Create a client object
	client := &Client{
		Conn:    conn,
		Message: make(chan *ws.WSMessage),
		RoomID:  roomID,
		UserID:  *models.CurrentUser,
	}

	h.Hub.Register <- client

	go client.writeMessage()
	go client.readMessage(h)
}

// Connect to chat notification channel
func (h *ChatHandler) NotificationHandler(c *gin.Context) {
	orgID := c.Param("id")
	// Connect to websocker
	conn, err := ws.WebsocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.WebSocketFailed)
		return
	}

	// Create SSE client and register to the notification hub
	client := &NotificationClient{
		Conn:         conn,
		OrgID:        orgID,
		UserID:       *models.CurrentUser,
		Notification: make(chan *MessageNotification),
	}

	h.NotificationHub.Register <- client

	go client.writeMessage(h.NotificationHub)
}

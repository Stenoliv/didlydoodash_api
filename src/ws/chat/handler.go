package chat

import (
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
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

func (h *ChatHandler) JoinRoom(c *gin.Context) {
	roomID := c.Param("chatId")
	// Retrive chat in database and check that user is part of
	room, err := daos.GetChat(roomID)
	if err != nil || room == nil {
		c.JSON(http.StatusForbidden, "Not in room"+err.Error())
		return
	}

	ok := false
	for _, member := range room.Members {
		if member.UserID == *models.CurrentUser {
			ok = true
		}
	}
	if !ok {
		c.JSON(http.StatusForbidden, "Not in room"+err.Error())
		return
	}

	// Create websocket connection
	conn, err := ws.WebsocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to initiate websocket connection")
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

func (h *ChatHandler) NotificationHandler(c *gin.Context) {
	// Connect to websocker
	conn, err := ws.WebsocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to initiate websocket connection")
		return
	}

	// Create SSE client and register to the notification hub
	client := &NotificationClient{
		Conn:         conn,
		UserID:       *models.CurrentUser,
		Notification: make(chan *MessageNotification),
	}

	h.NotificationHub.Register <- client

	go client.writeMessage(h.NotificationHub)
}

package chat

import (
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/ws"
	"encoding/json"
	"fmt"
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
	if err != nil {
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

	for _, msg := range room.Messages {
		data, err := json.Marshal(msg) // Marshal the message to JSON
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send previous messages"})
			return
		}

		client.Message <- &ws.WSMessage{
			Type:    utils.MessageSend,
			RoomID:  msg.RoomID,
			Payload: data,
		}

		notification := &MessageNotification{
			ChatID:  client.RoomID,
			Message: data,
		}
		h.NotificationHub.Broadcast <- notification
	}

	go client.writeMessage()
	go client.readMessage(h.Hub)
}

func (h *ChatHandler) NotificationHandler(c *gin.Context) {
	userID := c.Param("userID")

	// Set the necessary SSE headers
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	notificationChan := make(chan *MessageNotification)

	// Create SSE client and register to the notification hub
	client := &SSEClient{
		UserID:       userID,
		Notification: notificationChan,
	}

	h.NotificationHub.Register <- client

	defer func() {
		h.NotificationHub.Unregister <- client
	}()

	// Keep the connection open and listen for new notifications
	for {
		select {
		case notification := <-notificationChan:
			if _, err := fmt.Fprintf(c.Writer, "data: %s\n\n", *notification); err != nil {
				return // Client disconnected
			}
			c.Writer.Flush()

		case <-c.Writer.CloseNotify():
			return // Client disconnected
		}
	}
}

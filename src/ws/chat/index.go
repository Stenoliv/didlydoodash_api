package chat

import (
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/ws"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	Hub *Hub
}

func NewChatHandler() *ChatHandler {
	hub := NewHub()
	return &ChatHandler{Hub: hub}
}

func (h *ChatHandler) JoinRoom(c *gin.Context) {
	roomID := c.Param("roomId")
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
	}

	go client.writeMessage()
	go client.readMessage(h.Hub)
}

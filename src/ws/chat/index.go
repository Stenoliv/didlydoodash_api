package chat

import (
	"DidlyDoodash-api/src/db/daos"
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

type JoinRoomBody struct {
	RoomID string `json:"roomId" binding:"required"`
	UserID string `json:"userId" binding:"required"`
}

func (h *ChatHandler) JoinRoom(c *gin.Context) {
	var req JoinRoomBody

	// Bind the incoming request to validate user input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	// Retrive chat in database and check that user is part of
	room, err := daos.GetChatWithMessages(req.RoomID, req.UserID)
	if err != nil {
		c.JSON(http.StatusForbidden, "Not in room")
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
		RoomID:  req.RoomID,
		UserID:  req.UserID,
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
			RoomID:  msg.ChatRoomID,
			Payload: data,
		}
	}

	go client.writeMessage()
	go client.readMessage(h.Hub)
}

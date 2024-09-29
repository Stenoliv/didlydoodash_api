package ws

import (
	"DidlyDoodash-api/src/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

// Websocket upgrader
var WebsocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Base struct with type and json input
type WSMessage struct {
	Type    utils.MessageType `json:"type"`
	RoomID  string            `json:"roomId"`
	Payload json.RawMessage   `json:"payload"`
}

package whiteboardws

import (
	"DidlyDoodash-api/src/db/models"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Message chan models.LineData
	RoomID  string
	UserID  string
}

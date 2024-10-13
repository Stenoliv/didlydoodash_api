package whiteboardws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Message chan *WhiteboardMessage
	RoomID  string
	UserID  string
}

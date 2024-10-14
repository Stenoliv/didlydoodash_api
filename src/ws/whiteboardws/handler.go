package whiteboardws

import (
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/ws"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Wbhandler struct {
	Hub *Hub
}

func (wbh *Wbhandler) HandleConnections(w *gin.Context) {
	ws, err := ws.WebsocketUpgrader.Upgrade(w.Writer, w.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	Client := &Client{Conn: ws, UserID: *models.CurrentUser, RoomID: w.Param("wbID"), Message: make(chan *WhiteboardMessage)}
	wbh.Hub.Register <- Client
	fmt.Println("New client connected")
	go Client.readMessage(wbh)
	go Client.writeMessage()

}
func NewHandler() *Wbhandler {
	hub := NewHub()
	return &Wbhandler{Hub: hub}
}

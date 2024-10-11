package whiteboardws

import (
	"DidlyDoodash-api/src/db/models"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)
var clients = make(map[*websocket.Conn]bool)

var clientsMutex sync.Mutex

var broadcast = make(chan models.LineData)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
func HandleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error upgrading connection:", err)
        return
    }
    defer ws.Close()

    clientsMutex.Lock()
    clients[ws] = true
    clientsMutex.Unlock()
    fmt.Println("New client connected")

    for {
        var line models.LineData
        err := ws.ReadJSON(&line)
        if err != nil {
            fmt.Println("Error reading JSON:", err)

            clientsMutex.Lock()
            delete(clients, ws)
            clientsMutex.Unlock()
            break
        }

        fmt.Printf("Received line: %+v\n", line)

        broadcast <- line
    }
}
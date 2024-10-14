package whiteboardws

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
	"sync"
)

type WhiteboardMessage struct {
	RoomID  string   `json:"roomId" binding:"required"`
	Payload linedata `json:"payload" binding:"required"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *WhiteboardMessage
	mu         sync.Mutex
}

type Room struct {
	RoomID  string
	Clients map[string]*Client
	mu      sync.Mutex
}

type MessageWB struct {
	Data   *models.LineData `json:"linedata"`
	RoomId string           `json:"id"`
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			room, exists := h.Rooms[client.RoomID]
			if !exists {
				room = &Room{RoomID: client.RoomID, Clients: make(map[string]*Client)}
				h.Rooms[client.RoomID] = room
			}
			room.mu.Lock()
			room.Clients[client.UserID] = client
			room.mu.Unlock()
			h.mu.Unlock()
			var arr []models.LineData
			if err := db.DB.Model(&models.LineData{}).Where("whiteboard_id = ?", client.RoomID).Find(&arr).Error; err != nil {
				return
			}
			var points []float64
			for _, data := range arr {

				for _, point := range data.Points {
					points = append(points, point.Point)
				}
				message := &WhiteboardMessage{RoomID: client.RoomID, Payload: linedata{
					Stroke:      data.Stroke,
					StrokeWidth: data.StrokeWidth,
					Tool:        data.Tool,
					Text:        data.Text,
					Points:      points,
				}}
				h.Broadcast <- message
			}

		case client := <-h.Unregister:
			h.mu.Lock()
			if Room, exists := h.Rooms[client.RoomID]; exists {
				Room.mu.Lock()
				delete(Room.Clients, client.UserID)
				if len(Room.Clients) <= 0 {
					delete(h.Rooms, Room.RoomID)
				}
				Room.mu.Unlock()
			}
			h.mu.Unlock()
		case message := <-h.Broadcast:
			h.mu.Lock()
			if room, exists := h.Rooms[message.RoomID]; exists {
				room.mu.Lock()
				for _, client := range room.Clients {
					select {
					case client.Message <- message:
					default:
						close(client.Message)
						delete(room.Clients, client.UserID)
					}
				}
				room.mu.Unlock()
			}
			h.mu.Unlock()
		}

	}
}
func NewHub() *Hub {
	hub := &Hub{Rooms: map[string]*Room{}, Unregister: make(chan *Client), Register: make(chan *Client), Broadcast: make(chan *WhiteboardMessage)}
	go hub.run()
	return hub
}

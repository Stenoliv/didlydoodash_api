package chat

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/ws"
	"encoding/json"
	"sync"
)

type Room struct {
	ID      string             `json:"id"`
	Clients map[string]*Client `json:"clients"`
	mu      sync.Mutex
}

type Hub struct {
	Rooms      map[string]*Room
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *ws.WSMessage
	mu         sync.RWMutex
}

func NewHub() *Hub {
	hub := &Hub{
		Rooms:      make(map[string]*Room),
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *ws.WSMessage),
	}

	go hub.run()

	return hub
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client.UserID] = client

			// Check for room
			room, exists := h.Rooms[client.RoomID]
			if !exists {
				room = &Room{
					ID:      client.RoomID,
					Clients: make(map[string]*Client),
				}
				h.Rooms[client.RoomID] = room
			}

			// Add client to the room
			room.mu.Lock()
			room.Clients[client.UserID] = client
			room.mu.Unlock()

			h.mu.Unlock()

			// Get messages from the database and send if found
			var messages []models.ChatMessage
			// Load the first page (page 1 corresponds to offset 0)
			if err := db.DB.Model(&models.ChatMessage{}).
				Where("room_id = ?", client.RoomID).
				Order("created_at ASC").
				Scan(&messages).Error; err != nil {
				return
			}

			// Turn messages into json
			data, err := json.Marshal(messages)
			if err != nil {
				return
			}

			// Create response and send to client
			wsMessage := &ws.WSMessage{
				Type:    utils.SendMessages,
				RoomID:  client.RoomID,
				Payload: data,
			}
			client.Message <- wsMessage
		case client := <-h.Unregister:
			h.mu.Lock()
			if existingClient, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, existingClient.UserID)

				// Remove the client from the room if room exists
				if room, exists := h.Rooms[client.RoomID]; exists {
					room.mu.Lock()
					delete(room.Clients, client.UserID)

					if len(room.Clients) <= 0 {
						delete(h.Rooms, client.RoomID)
					}

					room.mu.Unlock()
				}
			}
			h.mu.Unlock()
		case message := <-h.Broadcast:
			h.mu.Lock()

			// Broadcast message to all clients in room
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

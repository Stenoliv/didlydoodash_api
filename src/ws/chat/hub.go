package chat

import (
	"DidlyDoodash-api/src/ws"
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
			room, exists := h.Rooms[client.RoomID]
			if exists {
				room.mu.Lock()
				room.Clients[client.UserID] = client
				room.mu.Unlock()
			}
			h.mu.Unlock()
		case client := <-h.Unregister:
			h.mu.Lock()
			if existingClient, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, existingClient.UserID)
				if room, exists := h.Rooms[client.RoomID]; exists {
					room.mu.Lock()
					delete(room.Clients, client.UserID)
					room.mu.Unlock()
				}
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

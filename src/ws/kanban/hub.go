package kanban

import (
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/utils"
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
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *ws.WSMessage
	mu         sync.RWMutex
}

// Function that returns new kanban hub
func NewHub() *Hub {
	hub := &Hub{
		Rooms:      map[string]*Room{},
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
			ClientRegister(h, client)
		case client := <-h.Unregister:
			ClientUnregister(h, client)
		case message := <-h.Broadcast:
			HandleBroadcast(h, message)
		}
	}
}

// Function to handle client register event
func ClientRegister(h *Hub, client *Client) {
	h.mu.Lock()
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

	// Get kanban user joined to
	kanban, err := daos.GetKanban(client.RoomID)
	if err != nil {
		client.SendErrorMessage("Failed to get kanban")
		return
	}

	archive, err := daos.GetKanbanArchive(client.RoomID)
	if err != nil {
		client.SendErrorMessage("Failed to get kanban archive")
		return
	}

	// Join message
	payload := &JoinMessage{
		Kanban:  *kanban,
		Archive: archive,
	}

	// Turn payload into json
	raw, err := payload.ToJSON()
	if err != nil {
		client.SendErrorMessage("Failed to get kanban")
		return
	}

	// Send join data to client
	client.Message <- &ws.WSMessage{
		Type:    utils.JoinKanban,
		RoomID:  client.RoomID,
		Payload: raw,
	}
}

// Function to handle client unregister event
func ClientUnregister(h *Hub, client *Client) {
	h.mu.Lock()

	// Remove the client from the room if room exists
	if room, exists := h.Rooms[client.RoomID]; exists {
		room.mu.Lock()
		delete(room.Clients, client.UserID)

		if len(room.Clients) <= 0 {
			delete(h.Rooms, client.RoomID)
		}

		room.mu.Unlock()
	}

	h.mu.Unlock()
}

// Function to handle broadcast events
func HandleBroadcast(h *Hub, message *ws.WSMessage) {
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

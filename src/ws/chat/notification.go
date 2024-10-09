package chat

import (
	"DidlyDoodash-api/src/db/daos"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type MessageNotification struct {
	ChatID  string          `json:"chatId"`
	Title   string          `json:"title"`
	Message json.RawMessage `json:"message"`
}

type NotificationClient struct {
	Conn         *websocket.Conn
	Notification chan *MessageNotification
	UserID       string
}

func (c *NotificationClient) writeMessage(hub *NotificationHub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Notification
		if !ok {
			break
		}

		if err := c.Conn.WriteJSON(message); err != nil {
			break
		}
	}
}

type NotificationHub struct {
	Clients    map[string]*NotificationClient
	Register   chan *NotificationClient
	Unregister chan *NotificationClient
	Broadcast  chan *MessageNotification
}

func NewChatHubSSE() *NotificationHub {
	hub := &NotificationHub{
		Clients:    make(map[string]*NotificationClient),
		Register:   make(chan *NotificationClient),
		Unregister: make(chan *NotificationClient),
		Broadcast:  make(chan *MessageNotification),
	}

	go hub.run()

	return hub
}

func (hub *NotificationHub) run() {
	for {
		select {
		case client := <-hub.Register:
			hub.Clients[client.UserID] = client
		case client := <-hub.Unregister:
			if _, ok := hub.Clients[client.UserID]; ok {
				delete(hub.Clients, client.UserID)
				close(client.Notification)
			}

		case message := <-hub.Broadcast:
			room, err := daos.GetChat(message.ChatID)
			if err != nil || room == nil {
				return
			}

			for _, member := range room.Members {
				if client, ok := hub.Clients[member.UserID]; ok {
					client.Notification <- message
				}
			}
		}
	}
}

func (hub *NotificationHub) NewNotification(chatID, title string, message json.RawMessage) {
	hub.Broadcast <- &MessageNotification{ChatID: chatID, Title: title, Message: message}
}

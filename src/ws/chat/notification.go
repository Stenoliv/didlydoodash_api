package chat

import (
	"encoding/json"
	"fmt"
)

type MessageNotification struct {
	ChatID  string          `json:"chatId"`
	Title   string          `json:"title"`
	Message json.RawMessage `json:"message"`
}

type SSEClient struct {
	UserID       string
	ChatID       string
	Notification chan *MessageNotification
}

type NotificationHub struct {
	Clients    map[string]*SSEClient
	Register   chan *SSEClient
	Unregister chan *SSEClient
	Broadcast  chan *MessageNotification
}

func NewChatHubSSE() *NotificationHub {
	hub := &NotificationHub{
		Clients:    make(map[string]*SSEClient),
		Register:   make(chan *SSEClient),
		Unregister: make(chan *SSEClient),
		Broadcast:  make(chan *MessageNotification),
	}

	go hub.Run()

	return hub
}

func (hub *NotificationHub) Run() {
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
			for _, client := range hub.Clients {
				client.Notification <- message
				fmt.Println(message)
			}
		}
	}
}

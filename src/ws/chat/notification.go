package chat

import (
	"DidlyDoodash-api/src/db/daos"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type MessageNotification struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type NotificationClient struct {
	Conn         *websocket.Conn
	Notification chan *MessageNotification
	OrgID        string
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

			// send to client number of new messages based on lastRead message
			go sendUnreadMessagesOnJoin(client)
		case client := <-hub.Unregister:
			if _, ok := hub.Clients[client.UserID]; ok {
				delete(hub.Clients, client.UserID)
				close(client.Notification)
			}

		case message := <-hub.Broadcast:
			switch message.Type {
			case MessageNew:
				var newMessage NewMessage
				if err := json.Unmarshal(message.Payload, &newMessage); err != nil {
					return
				}

				room, err := daos.GetChat(newMessage.ChatID)
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
}

func (hub *NotificationHub) NewNotification(notificationType string, payload json.RawMessage) {
	hub.Broadcast <- &MessageNotification{Type: notificationType, Payload: payload}
}

func sendUnreadMessagesOnJoin(c *NotificationClient) {
	org, err := daos.GetOrg(c.OrgID)
	if err != nil || org == nil {
		return
	}

	chats, err := daos.GetChats(org.ID)
	if err != nil {
		return
	}

	// For all chats user is part of send number of unread messages
	for _, chat := range chats {
		for _, member := range chat.Members {
			if member.UserID == c.UserID {
				num := member.GetNumOfUnreadMessage()

				// Create a count message data
				countMessage := &CountMessage{
					ChatID: chat.ID,
					UserID: c.UserID,
					Count:  num,
				}

				// Parse message to json
				payload, err := countMessage.ToJSON()
				if err != nil {
					return
				}

				// Create final notification for client
				notification := &MessageNotification{
					Type:    MessageCount,
					Payload: payload,
				}

				// Send notification to client
				c.Notification <- notification
			}
		}
	}
}

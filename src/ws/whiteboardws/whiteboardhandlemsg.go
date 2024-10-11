package whiteboardws

import (
	"fmt"
)

func HandleMessages() {
	for {
		line := <-broadcast
		for client := range clients {
			if err := client.WriteJSON(line); err != nil {
				fmt.Println("Error writing JSON:", err)
				client.Close() 
				delete(clients, client) 
			}
		}
	}
}

package whiteboardws

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Message chan *WhiteboardMessage
	RoomID  string
	UserID  string
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			break
		}

		if err := c.Conn.WriteJSON(message); err != nil {
			break
		}
	}
}

func (c *Client) readMessage(handler *Wbhandler) {
	defer func() {
		handler.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("unexpected error:", err)
			}
			break
		}
		var input WhiteboardMessage
		if err := json.Unmarshal(msg, &input); err != nil {
			fmt.Println("unmarshal error")
			return
		}
		tx := db.DB.Begin()
		data := &models.LineData{WhiteboardID: input.RoomID, Stroke: input.Payload.Stroke, StrokeWidth: input.Payload.StrokeWidth, Tool: input.Payload.Tool, Text: input.Payload.Text}
		if err := data.SaveLineData(tx); err != nil {
			fmt.Println("db begin payload save")
			tx.Rollback()
			return
		}
		for _, point := range input.Payload.Points {
			pt := models.LinePoint{Point: point, LineDataID: data.ID}
			if err := pt.SaveLinePoint(tx); err != nil {
				fmt.Println("error in looping of points")
				tx.Rollback()
				return
			}
			data.Points = append(data.Points, pt)
		}

		if err := tx.Commit().Error; err != nil {
			fmt.Println("commit error")
			tx.Rollback()
			return
		}
		fmt.Println(input)
		handler.Hub.Broadcast <- &input

	}
}

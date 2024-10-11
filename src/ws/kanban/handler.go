package kanban

import (
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Hub *Hub
}

func NewHandler() *Handler {
	hub := NewHub()
	return &Handler{Hub: hub}
}

func (h *Handler) JoinKanban(c *gin.Context) {
	projectID := c.Param("projectID")
	kanbanID := c.Param("kanbanID")

	project, err := daos.GetProject(projectID)
	if err != nil || project == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ProjectNotFound)
		return
	}

	conn, err := ws.WebsocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.WebSocketFailed)
		return
	}

	client := &Client{
		Conn:    conn,
		Message: make(chan *ws.WSMessage),
		RoomID:  kanbanID,
		UserID:  *models.CurrentUser,
	}

	h.Hub.Register <- client

	go client.writeMessage()
	go client.readMessage(h)
}

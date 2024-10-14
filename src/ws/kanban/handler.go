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

	member, err := daos.GetProjectMember(projectID, *models.CurrentUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, utils.ProjectMemberNotFound)
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
		Role:    member.Role,
	}

	h.Hub.Register <- client

	go client.writeMessage()
	go client.readMessage(h)
}

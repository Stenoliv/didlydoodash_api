package handlers

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/db/models"
	"DidlyDoodash-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChats(c *gin.Context) {
	orgID := c.Param("id")
	chats, err := daos.GetChats(orgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	c.JSON(http.StatusOK, gin.H{"chats": chats})
}

type CreateChatInput struct {
	Name    string             `json:"name" binding:"required"`
	Members []ChatMembersInput `json:"members" binding:"required"`
}

type ChatMembersInput struct {
	UserID string `json:"id" binding:"required"`
}

func CreateChat(c *gin.Context) {
	var req CreateChatInput
	orgID := c.Param("id")
	tx := db.DB.Begin()
	if err := c.ShouldBindJSON(&req); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	room := &models.ChatRoom{
		OrganisationID: orgID,
		Name:           req.Name,
	}

	if err := room.SaveChatRoom(tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	var members []models.ChatMember
	for _, member := range req.Members {
		roomMebmer := &models.ChatMember{
			UserID: member.UserID,
			RoomID: room.ID,
		}

		if err := roomMebmer.SaveMember(tx); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, utils.ServerError)
			return
		}

		members = append(members, *roomMebmer)
	}
	room.Members = members

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{"chat": room})
}

func AddUserToChat(c *gin.Context) {
	id := c.Param("id")
	chatId := c.Param("chatId")
	userId := c.Param("userId")
	tx := db.DB.Begin()

	member, err := daos.GetMember(id, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	if member == nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ServerError)
	}

	chatMember := &models.ChatMember{
		RoomID: chatId,
		UserID: userId,
	}

	if err := chatMember.SaveMember(tx); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"member": chatMember})
}

func RemoveUserToChat(c *gin.Context) {
	id := c.Param("id")
	chatId := c.Param("chatId")
	userId := c.Param("userId")

	if userId == *models.CurrentUser {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	member, err := daos.GetMember(id, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.InvalidInput)
		return
	}

	if member == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	chatMember, err := daos.GetChatMember(chatId, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	if err := db.DB.Delete(&chatMember).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"member": chatMember})
}

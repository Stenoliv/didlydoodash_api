package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
	"fmt"
)

func GetChats(orgID, userID string) ([]models.ChatRoom, error) {
	var chatRooms []models.ChatRoom

	err := db.DB.Model(&models.ChatRoom{}).
		Where("organisation_id = ?", orgID).
		Joins("JOIN organisation_schema.chat_members ON chat_rooms.id = organisation_schema.chat_members.chat_room_id").
		Where("organisation_schema.chat_members.user_id = ?", userID).
		Find(&chatRooms).Error

	if err != nil {
		return nil, err
	}

	return chatRooms, nil
}

func GetChatWithMessages(roomID, userID string) (room *models.ChatRoom, err error) {
	if err := db.DB.Model(&models.ChatRoom{}).Find(&room, roomID).Error; err != nil {
		return nil, err
	}

	if err := db.DB.Preload("Messages").Where("id = ?", roomID).First(room).Error; err != nil {
		return nil, err
	}

	if !isUserInRoom(userID, room.Members) {
		return nil, fmt.Errorf("user not in room: %s", userID)
	}

	return room, nil
}

func isUserInRoom(userID string, members []models.ChatMember) bool {
	for _, member := range members {
		if member.UserID == userID {
			return true
		}
	}
	return false
}

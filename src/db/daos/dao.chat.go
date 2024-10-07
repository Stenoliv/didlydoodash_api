package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func GetChats(orgID string) ([]models.ChatRoom, error) {
	var chatRooms []models.ChatRoom

	err := db.DB.Model(&models.ChatRoom{}).
		Where("organisation_id = ?", orgID).
		Joins("JOIN chat_members ON chat_members.room_id = chat_rooms.id AND chat_members.user_id = ?", models.CurrentUser).
		Find(&chatRooms).Error
	if err != nil {
		return nil, err
	}

	return chatRooms, nil
}

func GetChat(roomID string) (*models.ChatRoom, error) {
	var room models.ChatRoom

	err := db.DB.Model(&models.ChatRoom{}).
		Joins("JOIN chat_members ON chat_members.room_id = chat_rooms.id AND chat_members.user_id = ?", models.CurrentUser).
		Find(&room, "chat_rooms.id = ?", roomID).Error
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func GetChatMember(roomID, userID string) (*models.ChatMember, error) {
	var member models.ChatMember
	if err := db.DB.Model(&models.ChatMember{}).Where("room_id = ?", roomID).Where("user_id = ?", userID).Find(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

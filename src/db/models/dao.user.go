package models

import (
	"DidlyDoodash-api/src/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PublicUserData(db *gorm.DB) *gorm.DB {
	return db.Select("id", "username").Where("verified = true")
}

// Function to get all users with pagination enabled
func GetUsers(c *gin.Context) (PaginationResult[User], error) {
	var users []User

	// Use Paginate to fetch users and pagination metadata
	if err := db.DB.Scopes(Paginate[User](c)).Not("id = ?", CurrentUser).Find(&users).Error; err != nil {
		return PaginationResult[User]{}, err
	}

	// Retrieve pagination metadata directly from the Paginate function
	if paginated, exists := c.Get("pagination"); exists {
		pagination := paginated.(PaginationResult[User])
		return PaginationResult[User]{
			Total:      pagination.Total,
			Page:       pagination.Page,
			TotalPages: pagination.TotalPages,
			Data:       users,
		}, nil
	}

	// Return the users with an empty pagination if not found
	return PaginationResult[User]{Data: users}, nil
}

// Function to get user by id, username, email
func GetUser(id string) (User, error) {
	var user User
	if err := db.DB.Model(&User{}).Where("id = ? OR username = ? OR email = ?", id, id, id).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func GetSession(id string, jti string) (UserSession, error) {
	var session UserSession
	if err := db.DB.Model(&UserSession{}).Where("user_id = ? AND jti = ?", id, jti).First(&session).Error; err != nil {
		return UserSession{}, err
	}
	return session, nil
}

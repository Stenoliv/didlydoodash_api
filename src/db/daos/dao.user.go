package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PublicUserData(db *gorm.DB) *gorm.DB {
	return db.Select("id", "username").Where("verified = true")
}

// Function to get all users with pagination enabled
func GetUsers(c *gin.Context) (PaginationResult[models.User], error) {
	var users []models.User

	// Use Paginate to fetch users and pagination metadata
	if err := db.DB.Scopes(Paginate[models.User](c)).Not("id = ?", models.CurrentUser).Find(&users).Error; err != nil {
		return PaginationResult[models.User]{}, err
	}

	// Retrieve pagination metadata directly from the Paginate function
	if paginated, exists := c.Get("pagination"); exists {
		pagination := paginated.(PaginationResult[models.User])
		return PaginationResult[models.User]{
			Total:      pagination.Total,
			Page:       pagination.Page,
			TotalPages: pagination.TotalPages,
			Data:       users,
		}, nil
	}

	// Return the users with an empty pagination if not found
	return PaginationResult[models.User]{Data: users}, nil
}

// Function to get user by id, username, email
func GetUser(id string) (models.User, error) {
	var user models.User
	if err := db.DB.Model(&models.User{}).Where("id = ? OR username = ? OR email = ?", id, id, id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetSession(id string, jti string) (models.UserSession, error) {
	var session models.UserSession
	if err := db.DB.Model(&models.UserSession{}).Where("user_id = ? AND jti = ?", id, jti).First(&session).Error; err != nil {
		return models.UserSession{}, err
	}
	return session, nil
}

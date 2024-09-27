package daos

import (
	"DidlyDoodash-api/src/data"
	"DidlyDoodash-api/src/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PublicUserData(db *gorm.DB) *gorm.DB {
	return db.Select("id", "username").Where("verified = true")
}

// Function to get all users with pagination enabled
func GetUsers(c *gin.Context) (PaginationResult[data.User], error) {
	var users []data.User

	// Use Paginate to fetch users and pagination metadata
	if err := db.DB.Scopes(Paginate[data.User](c)).Not("id = ?", data.CurrentUser).Find(&users).Error; err != nil {
		return PaginationResult[data.User]{}, err
	}

	// Retrieve pagination metadata directly from the Paginate function
	if paginated, exists := c.Get("pagination"); exists {
		pagination := paginated.(PaginationResult[data.User])
		return PaginationResult[data.User]{
			Total:      pagination.Total,
			Page:       pagination.Page,
			TotalPages: pagination.TotalPages,
			Data:       users,
		}, nil
	}

	// Return the users with an empty pagination if not found
	return PaginationResult[data.User]{Data: users}, nil
}

// Function to get user by id, username, email
func GetUser(id string) (data.User, error) {
	var user data.User
	if err := db.DB.Model(&data.User{}).Where("id = ? OR username = ? OR email = ?", id, id, id).First(&user).Error; err != nil {
		return data.User{}, err
	}
	return user, nil
}

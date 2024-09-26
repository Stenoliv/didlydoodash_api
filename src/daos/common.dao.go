package daos

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaginationResult[T any] struct {
	Total      int `json:"total"`
	Page       int `json:"currentPage"`
	TotalPages int `json:"totalPages"`
	Data       []T `json:"data"`
}

// Checks for query page and limits and returns a paginated
func Paginate[T any](c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("limits", "10"))
		offset := (page - 1) * pageSize

		// Get total count
		var total int64
		db.Count(&total)

		// Calculate total pages
		totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

		// Store pagination metadata in the context
		c.Set("pagination", PaginationResult[T]{
			Total:      int(total),
			Page:       page,
			TotalPages: totalPages,
		})

		// Apply pagination
		return db.Offset(offset).Limit(pageSize)
	}
}

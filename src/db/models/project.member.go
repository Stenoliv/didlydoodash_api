package models

/**
 * Project members table
 */
type ProjectMembers struct {
	Project   Project `gorm:"" json:"-"`
	ProjectID string  `gorm:"size:21;" json:"-"`
	User      User    `gorm:"" json:"user"`
	UserID    string  `gorm:"size:21;" json:"-"`
}

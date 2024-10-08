package models

/**
 * Project members table
 */
type ProjectMembers struct {
	ProjectID string  `gorm:"size:21;" json:"-"`
	Project   Project `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserID    string  `gorm:"size:21;" json:"-"`
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

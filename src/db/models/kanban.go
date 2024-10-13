package models

import (
	"DidlyDoodash-api/src/db/datatypes"
	"time"

	"gorm.io/gorm"
)

// Kanban struct
type Kanban struct {
	Base
	ProjectID  string                 `gorm:"size:21;not null;uniqueIndex:idx_kanban;" json:"projectId"`
	Project    Project                `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Name       string                 `gorm:"size:50;not null;uniqueIndex:idx_kanban;" json:"name"`
	Status     datatypes.KanbanStatus `gorm:"type:kanban_status;not null;default:Planning;" json:"status"`
	Categories []KanbanCategory       `gorm:"" json:"categories"`
}

func (k *Kanban) BeforeCreate(tx *gorm.DB) error {
	if err := k.GenerateID(); err != nil {
		return err
	}
	return nil
}

func (k *Kanban) AfterCreate(tx *gorm.DB) error {
	if err := tx.Model(&KanbanCategory{}).Where("kanban_id = ?", k.ID).Order("created_at ASC").Find(&k.Categories).Error; err != nil {
		return err
	}
	return nil
}

func (k *Kanban) AfterFind(tx *gorm.DB) error {
	if err := tx.Model(&KanbanCategory{}).Where("kanban_id = ?", k.ID).Order("created_at ASC").Find(&k.Categories).Error; err != nil {
		return err
	}
	return nil
}

func (k *Kanban) SaveKanban(tx *gorm.DB) error {
	return tx.Create(&k).Error
}

// Kanban Category struct
type KanbanCategory struct {
	Base
	DeletedAt gorm.DeletedAt `gorm:"" json:"deletedAt"`
	KanbanID  string         `gorm:"size:21;not null;" json:"kanbanId"`
	Name      string         `gorm:"size:50;" json:"name"`
	Items     []KanbanItem   `gorm:"" json:"items"`
}

func (k *KanbanCategory) BeforeCreate(tx *gorm.DB) error {
	if err := k.GenerateID(); err != nil {
		return err
	}
	return nil
}

func (k *KanbanCategory) AfterFind(tx *gorm.DB) error {
	if err := tx.Model(&KanbanItem{}).Where("kanban_category_id = ?", k.ID).Order("created_at ASC").Find(&k.Items).Error; err != nil {
		return err
	}
	return nil
}

func (k *KanbanCategory) SaveCategory(tx *gorm.DB) error {
	return tx.Create(&k).Error
}

// Kanban Item struct
type KanbanItem struct {
	Base
	KanbanCategoryID string                       `gorm:"size:21;not nill;" json:"categoryId"`
	DeletedAt        gorm.DeletedAt               `gorm:"" json:"deletedAt"`
	DueDate          *time.Time                   `gorm:"null" json:"due_date"`
	EstimatedTime    *int                         `gorm:"null" json:"estimated_time"`
	Priority         datatypes.KanbanItemPriority `gorm:"type:kanban_item_priority;default:None;not null;" json:"priority"`
	Title            string                       `gorm:"size:40;not null;" json:"title"`
	Description      string                       `gorm:"" json:"description"`
	AssignedID       *string                      `gorm:"size:21;null" json:"-"`
	Assigned         User                         `gorm:"constraint:OnUpdate:CASCADE;OnDelete:CASCADE;" json:"assigned"`
}

func (k *KanbanItem) BeforeCreate(tx *gorm.DB) error {
	if err := k.GenerateID(); err != nil {
		return err
	}
	return nil
}

func (k *KanbanItem) SaveItem(tx *gorm.DB) error {
	return tx.Create(&k).Error
}

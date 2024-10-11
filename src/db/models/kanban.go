package models

import (
	"DidlyDoodash-api/src/db/datatypes"

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
	if err := tx.Model(&KanbanCategory{}).Where("kanban_id = ?", k.ID).Find(&k.Categories).Error; err != nil {
		return err
	}
	return nil
}

func (k *Kanban) AfterFind(tx *gorm.DB) error {
	if err := tx.Model(&KanbanCategory{}).Where("kanban_id = ?", k.ID).Find(&k.Categories).Error; err != nil {
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
	KanbanID string       `gorm:"size:21;not null;" json:"kanbanId"`
	Name     string       `gorm:"size:50;" json:"name"`
	Items    []KanbanItem `gorm:"" json:"items"`
}

func (k *KanbanCategory) BeforeCreate(tx *gorm.DB) error {
	if err := k.GenerateID(); err != nil {
		return err
	}
	return nil
}

func (k *KanbanCategory) AfterFind(tx *gorm.DB) error {
	if err := tx.Model(&KanbanItem{}).Where("kanban_category_id = ?", k.ID).Find(&k.Items).Error; err != nil {
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
	KanbanCategoryID string `gorm:"size:21;not nill;" json:"categoryId"`
	Title            string `gorm:"size:20;not null;" json:"title"`
	Desc             string `gorm:"" json:"desc"`
}

func (k *KanbanItem) BeforeCreate(tx *gorm.DB) error {
	if err := k.GenerateID(); err != nil {
		return err
	}
	return nil
}

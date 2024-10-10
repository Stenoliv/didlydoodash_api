package models

import "gorm.io/gorm"

// Kanban struct
type Kanban struct {
	Base
	ProjectID  string           `gorm:"size:21;not null;uniqueIndex:idx_kanban;" json:"projectId"`
	Project    Project          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Name       string           `gorm:"size:50;not null;uniqueIndex:idx_kanban;" json:"name"`
	Categories []KanbanCategory `gorm:"-" json:"categories"`
}

func (k *Kanban) BeforeCreate(tx *gorm.DB) error {
	if err := k.GenerateID(); err != nil {
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
	Name  string       `gorm:"size:50;" json:"name"`
	Items []KanbanItem `gorm:"-" json:"items"`
}

func (k *KanbanCategory) BeforeCreate(tx *gorm.DB) error {
	if err := k.GenerateID(); err != nil {
		return err
	}
	return nil
}

// Kanban Item struct
type KanbanItem struct {
	Base
	CategoryID string         `gorm:"size:21;not nill;" json:"categoryId"`
	Category   KanbanCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (k *KanbanItem) BeforeCreate(tx *gorm.DB) error {
	if err := k.GenerateID(); err != nil {
		return err
	}
	return nil
}

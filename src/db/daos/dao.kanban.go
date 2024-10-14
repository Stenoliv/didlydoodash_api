package daos

import (
	"DidlyDoodash-api/src/db"
	"DidlyDoodash-api/src/db/models"
)

func GetAllKanbans(id string) (kanbans []models.Kanban, err error) {
	if err := db.DB.Where("project_id = ?", id).Order("created_at ASC").Find(&kanbans).Error; err != nil {
		return nil, err
	}
	return kanbans, nil
}

func GetKanban(id string) (*models.Kanban, error) {
	var kanban *models.Kanban
	if err := db.DB.Where("id = ?", id).First(&kanban).Error; err != nil {
		return nil, err
	}
	return kanban, nil
}

func GetCategory(id string) (*models.KanbanCategory, error) {
	var category *models.KanbanCategory
	if err := db.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func GetDeletedCategory(id string) (*models.KanbanCategory, error) {
	var category *models.KanbanCategory
	if err := db.DB.Unscoped().Where("id = ? AND deleted_at IS NOT NULL", id).First(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func GetKanbanItem(id string) (*models.KanbanItem, error) {
	var item *models.KanbanItem
	if err := db.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func GetDeletedItem(id string) (*models.KanbanItem, error) {
	var item *models.KanbanItem
	if err := db.DB.Unscoped().Where("id = ? AND deleted_at IS NOT NULL", id).First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func GetKanbanArchive(kanbanID string) ([]models.KanbanArchiveItem, error) {
	var archiveItems []models.KanbanArchiveItem

	// Get all soft-deleted categories associated with the kanbanID
	var categories []models.KanbanCategory
	if err := db.DB.Unscoped().Model(&categories).Select("id, name, deleted_at").Where("kanban_id = ? AND deleted_at IS NOT NULL", kanbanID).Scan(&categories).Error; err != nil {
		return nil, err
	}

	// Add soft-deleted categories to the archiveItems slice
	for _, category := range categories {
		archiveItems = append(archiveItems, models.KanbanArchiveItem{
			ID:        category.ID,
			Type:      "category",
			DeletedAt: category.DeletedAt,
			Name:      category.Name,
		})
	}

	// Get all soft-deleted items associated with the kanbanID by joining with categories
	var items []models.KanbanItem
	if err := db.DB.Unscoped().Model(&items).Select("kanban_items.id, kanban_items.title, kanban_items.deleted_at, kanban_items.kanban_category_id").
		Joins("JOIN kanban_categories ON kanban_items.kanban_category_id = kanban_categories.id").
		Where("kanban_categories.kanban_id = ? AND kanban_items.deleted_at IS NOT NULL", kanbanID).
		Scan(&items).Error; err != nil {
		return nil, err
	}

	// Add soft-deleted items to the archiveItems slice
	for _, item := range items {
		archiveItems = append(archiveItems, models.KanbanArchiveItem{
			ID:        item.ID,
			Type:      "item",
			DeletedAt: item.DeletedAt,
			Name:      item.Title,
		})
	}

	return archiveItems, nil
}

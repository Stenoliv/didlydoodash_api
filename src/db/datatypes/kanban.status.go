package datatypes

import (
	"fmt"
	"strings"
)

const (
	KanbanStatusName string = "kanban_status"
)

// Kanban Status
type KanbanStatus string

const (
	KANBAN_PLANNING    KanbanStatus = "Planning"
	KANBAN_IN_PROGRESS KanbanStatus = "In Progress"
	KANBAN_DONE        KanbanStatus = "Done"
	KANBAN_ARCHIVED    KanbanStatus = "Archived"
)

var KanbanStatusEnum = []KanbanStatus{
	KANBAN_PLANNING,
	KANBAN_IN_PROGRESS,
	KANBAN_DONE,
	KANBAN_ARCHIVED,
}

// GetKanbanStatusEnum returns a string representation of the valid Kanban statuses
func GetKanbanStatusEnum(statuses []KanbanStatus) string {
	var statusStrings []string
	for _, status := range statuses {
		statusStrings = append(statusStrings, fmt.Sprintf("'%s'", status))
	}
	return strings.Join(statusStrings, ", ")
}

// IsValidKanbanStatus checks if the provided status is a valid Kanban status
func IsValidKanbanStatus(status string) bool {
	for _, validStatus := range KanbanStatusEnum {
		if strings.EqualFold(string(validStatus), status) {
			return true
		}
	}
	return false
}

// ToKanbanStatus converts a string to a KanbanStatus type, returning PLANNING if invalid
func ToKanbanStatus(status string) KanbanStatus {
	for _, validStatus := range KanbanStatusEnum {
		if strings.EqualFold(string(validStatus), status) {
			return validStatus
		}
	}
	return KANBAN_PLANNING
}

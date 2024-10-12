package datatypes

import (
	"fmt"
	"strings"
)

const (
	KanbanItemPriorityName string = "kanban_item_priority"
)

type KanbanItemPriority string

const (
	// Executive & Leadership Roles
	EXTREME KanbanItemPriority = "Extreme"
	HIGH    KanbanItemPriority = "High"
	MEDIUM  KanbanItemPriority = "Medium"
	LOW     KanbanItemPriority = "Low"
	NONE    KanbanItemPriority = "None"
)

var KanbanItemPriorityEnum = []KanbanItemPriority{
	EXTREME,
	HIGH,
	MEDIUM,
	LOW,
	NONE,
}

func GetKanbanItemPriorityEnum(roles []KanbanItemPriority) string {
	var roleStrings []string
	for _, role := range roles {
		roleStrings = append(roleStrings, fmt.Sprintf("'%s'", role))
	}
	return strings.Join(roleStrings, ", ")
}

func IsValidKanbanItemPriority(role string) bool {
	for _, validRole := range KanbanItemPriorityEnum {
		if strings.EqualFold(string(validRole), role) {
			return true
		}
	}
	return false
}

func ToKanbanItemPriority(role string) KanbanItemPriority {
	for _, validRole := range KanbanItemPriorityEnum {
		if strings.EqualFold(string(validRole), role) {
			return validRole
		}
	}
	return NONE
}

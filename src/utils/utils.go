package utils

import (
	"DidlyDoodash-api/src/db"
	"reflect"
	"strings"
)

func GetTableName(schema string, model interface{}) string {
	typeName := reflect.TypeOf(model).Elem().Name()
	tableName := ToSnakeCase(typeName)
	return schema + db.DB.NamingStrategy.TableName(tableName)
}

func ToSnakeCase(s string) string {
	var result strings.Builder

	for i, runeValue := range s {
		if i > 0 && i < len(s)-1 && 'A' <= runeValue && runeValue <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(runeValue)
	}

	return strings.ToLower(result.String())
}
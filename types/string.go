package types

import (
	"strings"
)

// IsString returns true if underlying given data type
// can hold or produce string values.
func IsString(v interface{}) bool {
	if v == nil {
		return false
	}

	switch x := v.(type) {
	case string:
		return true
	case *string:
		return true
	case anyDataTypeProvider:
		return IsString(x.DataType())
	case databaseTypeNameDescriber:
		return isDatabaseString(x.DatabaseTypeName())
	}

	return false
}

func isDatabaseString(v string) bool {
	switch strings.ToLower(v) {
	case "varchar", "char":
		return true
	case "enum":
		return true
	}
	return false
}

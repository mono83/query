package types

import (
	"strings"
)

// IsInt returns true if underlying given data type
// can hold or produce integer values.
func IsInt(v interface{}) bool {
	if v == nil {
		return false
	}

	switch x := v.(type) {
	case int, int8, int16, int32, int64:
		return true
	case uint, uint8, uint16, uint32, uint64:
		return true
	case *int, *int8, *int16, *int32, *int64:
		return true
	case *uint, *uint8, *uint16, *uint32, *uint64:
		return true
	case anyDataTypeProvider:
		return IsInt(x.DataType())
	case databaseTypeNameDescriber:
		return isDatabaseInt(x.DatabaseTypeName())
	}

	return false
}

func isDatabaseInt(v string) bool {
	switch strings.ToLower(v) {
	case "tiny", "mediumint", "int", "bigint":
		return true
	}
	return false
}

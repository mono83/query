package types

import "strings"

// IsFloat returns true if underlying given data type
// can hold or produce float values.
func IsFloat(v interface{}) bool {
	if v == nil {
		return false
	}

	switch x := v.(type) {
	case float32, float64:
		return true
	case *float32, *float64:
		return true
	case anyDataTypeProvider:
		return IsFloat(x.DataType())
	case databaseTypeNameDescriber:
		return isDatabaseFloat(x.DatabaseTypeName())
	}

	return false
}

func isDatabaseFloat(v string) bool {
	switch strings.ToLower(v) {
	case "float", "decimal":
		return true
	}
	return false
}

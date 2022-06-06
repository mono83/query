package types

import (
	"reflect"
)

// IsNullable return true if type of given value supports
// nil as valid value.
func IsNullable(v interface{}) bool {
	if v == nil {
		return false
	}

	switch x := v.(type) {
	case anyDataTypeProvider:
		return IsNullable(x.DataType())
	case nullableDescriber:
		y, z := x.Nullable()
		return y && z
	case reflect.Type:
		return IsNullable(x.Kind())
	case reflect.Kind:
		return x == reflect.Ptr || x == reflect.Slice || x == reflect.Map || x == reflect.Chan || x == reflect.Func
	}

	return IsNullable(reflect.TypeOf(v))
}

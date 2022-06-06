package types

// IsNumber returns true if underlying given data type
// can hold or produce any type of number values.
func IsNumber(v interface{}) bool {
	return IsInt(v) || IsFloat(v)
}

package eval

func is(v interface{}) (isSlice, isBool, isString, isInt, isFloat bool) {
	if v != nil {
		switch v.(type) {
		case bool:
			isBool = true
		case []bool:
			isSlice = true
			isBool = true
		case string:
			isString = true
		case []string:
			isSlice = true
			isString = true
		case float32, float64:
			isFloat = true
		case []float32, []float64:
			isSlice = true
			isFloat = true
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			isInt = true
		case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
			isSlice = true
			isInt = true
		case []interface{}:
			isSlice = true
		}
	}

	return
}

func numberInt64(v interface{}) (int64, bool) {
	switch x := v.(type) {
	case int:
		return int64(x), true
	case int64:
		return x, true
	case int32:
		return int64(x), true
	case int16:
		return int64(x), true
	case int8:
		return int64(x), true
	case uint64:
		return int64(x), true
	case uint32:
		return int64(x), true
	case uint16:
		return int64(x), true
	case uint8:
		return int64(x), true
	case float64:
		return int64(x), true
	case float32:
		return int64(x), true
	}

	return 0, false
}

func numberFloat64(v interface{}) (float64, bool) {
	switch x := v.(type) {
	case int:
		return float64(x), true
	case int64:
		return float64(x), true
	case int32:
		return float64(x), true
	case int16:
		return float64(x), true
	case int8:
		return float64(x), true
	case uint64:
		return float64(x), true
	case uint32:
		return float64(x), true
	case uint16:
		return float64(x), true
	case uint8:
		return float64(x), true
	case float64:
		return x, true
	case float32:
		return float64(x), true
	}

	return 0, false
}

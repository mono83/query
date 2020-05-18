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
	switch v.(type) {
	case int:
		return int64(v.(int)), true
	case int64:
		return v.(int64), true
	case int32:
		return int64(v.(int32)), true
	case int16:
		return int64(v.(int16)), true
	case int8:
		return int64(v.(int8)), true
	case uint64:
		return int64(v.(uint64)), true
	case uint32:
		return int64(v.(uint32)), true
	case uint16:
		return int64(v.(uint16)), true
	case uint8:
		return int64(v.(uint8)), true
	case float64:
		return int64(v.(float64)), true
	case float32:
		return int64(v.(float32)), true
	}

	return 0, false
}

func numberFloat64(v interface{}) (float64, bool) {
	switch v.(type) {
	case int:
		return float64(v.(int)), true
	case int64:
		return float64(v.(int64)), true
	case int32:
		return float64(v.(int32)), true
	case int16:
		return float64(v.(int16)), true
	case int8:
		return float64(v.(int8)), true
	case uint64:
		return float64(v.(uint64)), true
	case uint32:
		return float64(v.(uint32)), true
	case uint16:
		return float64(v.(uint16)), true
	case uint8:
		return float64(v.(uint8)), true
	case float64:
		return v.(float64), true
	case float32:
		return float64(v.(float32)), true
	}

	return 0, false
}

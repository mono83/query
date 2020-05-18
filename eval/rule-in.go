package eval

import "reflect"

func (e Evaluator) in(left, right interface{}) bool {
	isSlice, _, _, _, _ := is(right)
	if !isSlice {
		return e.equals(left, right)
	}

	s := reflect.ValueOf(right)
	for i := 0; i < s.Len(); i++ {
		if e.equals(left, s.Index(i).Interface()) {
			// At least one value in slice equals
			return true
		}
	}

	return false
}

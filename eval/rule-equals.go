package eval

import "fmt"

func (e Evaluator) equals(left, right interface{}) bool {
	if left == right {
		return true
	}
	if e.Strict {
		// Strict mode
		return false
	}

	return fmt.Sprint(left) == fmt.Sprint(right)
}

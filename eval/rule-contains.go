package eval

import (
	"fmt"
	"strings"
)

func (e Evaluator) contains(left, right interface{}) bool {
	if left == nil || right == nil {
		return false
	}
	sleft, okleft := left.(string)
	sright, okright := right.(string)
	if e.Strict && (!okleft || !okright) {
		// Strict mode
		return false
	}
	if !okleft {
		sleft = fmt.Sprint(left)
	}
	if !okright {
		sright = fmt.Sprint(right)
	}

	return strings.Contains(sleft, sright)
}

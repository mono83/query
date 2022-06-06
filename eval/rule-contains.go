package eval

import (
	"github.com/mono83/query/types"
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
		sleft = types.ToString(left)
	}
	if !okright {
		sright = types.ToString(right)
	}

	return strings.Contains(sleft, sright)
}

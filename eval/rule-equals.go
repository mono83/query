package eval

import (
	"github.com/mono83/query/types"
)

func (e Evaluator) equals(left, right interface{}) bool {
	if left == right {
		return true
	}
	if e.Strict {
		// Strict mode
		return false
	}

	return types.ToString(left) == types.ToString(right)
}

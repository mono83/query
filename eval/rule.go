package eval

import (
	"github.com/mono83/query"
	"github.com/mono83/query/match"
)

// Rule evaluates given rule
func (e Evaluator) Rule(r query.Rule) bool {
	if r == nil {
		return false
	}
	op := r.Type()
	left := r.Left()
	right := r.Right()

	if op == match.Unknown || op.IsCustom() {
		// Unknown or custom operator
		return false
	}

	switch op {
	case match.IsNull:
		return right == nil
	case match.NotIsNull:
		return right != nil
	case match.Equals:
		return e.equals(left, right)
	case match.NotEquals:
		return !e.equals(left, right)
	case match.In:
		return e.in(left, right)
	case match.NotIn:
		return !e.in(left, right)
	case match.GreaterThan, match.GreaterThanEquals, match.LesserThan, match.LesserThanEquals:
		return e.comp(left, right, op)
	}

	return false
}

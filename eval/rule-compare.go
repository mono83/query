package eval

import (
	"reflect"

	"github.com/mono83/query/match"
)

func (e Evaluator) comp(left, right interface{}, op match.Type) bool {
	if e.Strict {
		if reflect.TypeOf(left) != reflect.TypeOf(right) {
			return false
		}
	}

	_, _, _, lInt, lFloat := is(left)
	_, _, _, rInt, rFloat := is(right)
	if lFloat || rFloat {
		l64, lok := numberFloat64(left)
		r64, rok := numberFloat64(right)
		if !lok || !rok {
			return false
		}
		return e.compFloat64(l64, r64, op)
	}
	if lInt || rInt {
		l64, lok := numberInt64(left)
		r64, rok := numberInt64(right)
		if !lok || !rok {
			return false
		}
		return e.compInt64(l64, r64, op)
	}

	return false
}

func (e Evaluator) compFloat64(left, right float64, op match.Type) bool {
	switch op {
	case match.GreaterThan:
		return left > right
	case match.GreaterThanEquals:
		return left >= right
	case match.LesserThan:
		return left < right
	case match.LesserThanEquals:
		return left <= right
	}

	return false
}

func (e Evaluator) compInt64(left, right int64, op match.Type) bool {
	switch op {
	case match.GreaterThan:
		return left > right
	case match.GreaterThanEquals:
		return left >= right
	case match.LesserThan:
		return left < right
	case match.LesserThanEquals:
		return left <= right
	}

	return false
}

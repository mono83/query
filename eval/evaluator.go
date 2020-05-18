package eval

import "github.com/mono83/query"

// Evaluator provides logic to evaluate rules and conditions
type Evaluator struct {
	Strict bool
}

// ConditionStrict evaluates condition in strict mode
func ConditionStrict(c query.Condition) bool {
	return Evaluator{Strict: true}.Condition(c)
}

// Condition evaluates condition in non-strict mode
func Condition(c query.Condition) bool {
	return Evaluator{Strict: false}.Condition(c)
}

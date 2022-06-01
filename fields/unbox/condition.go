package unbox

import (
	"github.com/mono83/query"
	"github.com/mono83/query/conditions"
	"github.com/mono83/query/rules"
)

// Condition iterates over all rules (including nested) and
// replaces all found fields with their string name by calling
// field.Name method.
func Condition(c query.Condition) query.Condition {
	return conditions.Map(c, func(rule query.Rule) query.Rule {
		left, right := rule.Left(), rule.Right()
		if f, ok := left.(query.Field); ok {
			left = f.Name()
		}
		if f, ok := right.(query.Field); ok {
			right = f.Name()
		}
		return rules.New(left, rule.Type(), right)
	})
}

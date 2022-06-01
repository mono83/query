package box

import (
	"fmt"
	"github.com/mono83/query"
	"github.com/mono83/query/conditions"
	"github.com/mono83/query/rules"
)

// ConditionLeft iterates over all rules (even nested) and replaces
// all string field names (left part of rule) with their query.Field
// definitions.
// Method will fail on unknown fields.
func ConditionLeft(c query.Condition, fs []query.Field) (query.Condition, error) {
	m := fieldNameMap(fs)
	var err error
	boxed := conditions.Map(c, func(rule query.Rule) query.Rule {
		if err == nil {
			left := rule.Left()
			name := ""
			if x, ok := left.(string); ok {
				name = x
			} else if x, ok := left.(query.Named); ok {
				name = x.Name()
			} else {
				err = fmt.Errorf("unable to resolve left in rule %s", rule)
			}

			if len(name) > 0 {
				if f, ok := m[name]; ok {
					return rules.WithLeft(rule, f)
				}
				err = fmt.Errorf("field named %s not found for rule %s", name, rule)
			}
		}
		return rule
	})
	if err != nil {
		return nil, err
	}

	return boxed, nil
}

package eval

import "github.com/mono83/query"

// Condition evaluates given condition
func (e Evaluator) Condition(c query.Condition) bool {
	if c == nil {
		return false
	}

	logic := c.GetType()
	rules := c.GetRules()
	conditions := c.GetConditions()

	if len(rules) == 0 && len(conditions) == 0 {
		return true
	}

	for _, rule := range rules {
		if e.Rule(rule) {
			// TRUE
			if logic == query.Or {
				// At least one rule is true
				return true
			}
		} else {
			// FALSE
			if logic == query.And {
				// At least one rule is false
				return false
			}
		}
	}

	for _, condition := range conditions {
		if e.Condition(condition) {
			// TRUE
			if logic == query.Or {
				// At least one condition is true
				return true
			}
		} else {
			// FALSE
			if logic == query.And {
				// At least one condition is false
				return false
			}
		}
	}

	if logic == query.Or {
		return false
	}

	return true
}

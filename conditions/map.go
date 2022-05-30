package conditions

import "github.com/mono83/query"

// Map applies mapping function to all rules (including nested)
// within given condition.
func Map(c query.Condition, f func(query.Rule) query.Rule) query.Condition {
	if f == nil || IsEmpty(c) {
		return c
	}

	var rules []query.Rule
	for _, x := range c.Rules() {
		mapped := f(x)
		if mapped != nil {
			rules = append(rules, mapped)
		}
	}

	var conditions []query.Condition
	for _, x := range c.Conditions() {
		mapped := Map(x, f)
		if mapped != nil {
			conditions = append(conditions, mapped)
		}
	}

	return New(c.Type(), rules, conditions)
}

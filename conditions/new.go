package conditions

import "github.com/mono83/query"

// New builds and returns new condition
func New(logic query.Logic, rules []query.Rule, conditions []query.Condition) query.Condition {
	if logic == query.And && len(conditions) == 0 {
		return commonCondition(rules)
	}
	return condition{
		logic:      logic,
		rules:      rules,
		conditions: conditions,
	}
}

// ForAllRules returns conditions built for all rules with logic AND
func ForAllRules(rules ...query.Rule) query.Condition {
	return New(query.And, rules, nil)
}

// ForAnyRule returns conditions built for all rules with logic OR
func ForAnyRule(rules ...query.Rule) query.Condition {
	return New(query.Or, rules, nil)
}

type condition struct {
	logic      query.Logic
	rules      []query.Rule
	conditions []query.Condition
}

func (c condition) Type() query.Logic             { return c.logic }
func (c condition) Rules() []query.Rule           { return c.rules }
func (c condition) Conditions() []query.Condition { return c.conditions }

// commonCondition is condition with AND logic and without nested conditions
type commonCondition []query.Rule

func (c commonCondition) Type() query.Logic             { return query.And }
func (c commonCondition) Rules() []query.Rule           { return []query.Rule(c) }
func (c commonCondition) Conditions() []query.Condition { return nil }

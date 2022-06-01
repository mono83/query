package conditions

import "github.com/mono83/query"

// VisitRules applies callback function to every rule within
// condition and nested conditions.
func VisitRules(c query.Condition, f func(query.Rule)) {
	if f != nil {
		VisitRulesUntil(c, func(rule query.Rule) bool {
			f(rule)
			return true
		})
	}
}

// VisitRulesUntil applies callback function to every rule within
// condition and nested conditions until at least callback function
// returns false.
// VisitRulesUntil returns true only if given callback function is
// not nil and all invocations of callback function (if any) also
// returns true.
func VisitRulesUntil(c query.Condition, f func(query.Rule) bool) bool {
	if f == nil {
		return false
	}
	if IsEmpty(c) {
		return true
	}

	for _, r := range c.Rules() {
		if success := f(r); !success {
			return false
		}
	}

	for _, c := range c.Conditions() {
		if success := VisitRulesUntil(c, f); !success {
			return false
		}
	}
	return true
}

// CountRules returns amount of rules without condition
func CountRules(c query.Condition) (out int) {
	VisitRules(c, func(query.Rule) { out++ })
	return
}

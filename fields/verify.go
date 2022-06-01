package fields

import (
	"fmt"
	"github.com/mono83/query"
	"github.com/mono83/query/conditions"
)

// VerifyCondition checks that condition has only fields on
// the left side of rules inside it and those fields are
// allowed to be used as filter (i.e. in WHERE clause).
//
// This method will fail if string names are used instead
// of fields - to fix it apply boxing first.
func VerifyCondition(c query.Condition) (err error) {
	conditions.VisitRulesUntil(c, func(rule query.Rule) bool {
		left := rule.Left()
		if f, ok := left.(query.Field); ok {
			if !f.Filterable() {
				err = fmt.Errorf(`field "%s" used in %s does not support filtering`, f.Name(), rule)
				return false
			}
		} else {
			err = fmt.Errorf("no field found in rule %s", rule)
			return false
		}
		return true
	})
	return
}

// VerifyFilter checks that condition has only fields on
// the left side of rules inside it and those fields are
// allowed to be used as filter (i.e. in WHERE clause).
//
// Also, it checks that sorting is using fields allowing
// sorting operation too.
//
// This method will fail if string names are used instead
// of fields - to fix it apply boxing first.
func VerifyFilter(c query.Filter) (err error) {
	err = VerifyCondition(c)
	if err == nil {
		for _, s := range c.Sorting() {
			if f, ok := s.(query.Field); ok {
				if !f.Sortable() {
					err = fmt.Errorf(`field "%s" does not support sorting`, f.Name())
					break
				}
			} else {
				err = fmt.Errorf(`no field found in sorting by "%s"`, s)
			}
		}
	}
	return
}

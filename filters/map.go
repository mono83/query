package filters

import (
	"github.com/mono83/query"
	"github.com/mono83/query/conditions"
)

// Map applies mapping function to all rules (including nested)
// within given filter.
func Map(f query.Filter, mf func(query.Rule) query.Rule) query.Filter {
	if f == nil || conditions.IsEmpty(f) {
		return f
	}

	return FromCondition(conditions.Map(f, mf), f.Sorting(), f.Limit(), f.Offset())
}

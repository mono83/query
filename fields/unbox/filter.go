package unbox

import (
	"github.com/mono83/query"
	"github.com/mono83/query/filters"
	"github.com/mono83/query/sorting"
)

// Filter iterates over sorting settings and rules (including
// nested) and replaces all found fields with their string
// name by calling field.Name method.
func Filter(c query.Filter) query.Filter {
	cond := Condition(c)
	var srt []query.Sorting
	for _, s := range c.Sorting() {
		if f, ok := s.(query.Field); ok {
			srt = append(srt, sorting.String(f.Name(), s.Type()))
		} else {
			srt = append(srt, s)
		}
	}
	return filters.FromCondition(cond, srt, c.Limit(), c.Offset())
}

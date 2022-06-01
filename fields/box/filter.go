package box

import (
	"fmt"
	"github.com/mono83/query"
	"github.com/mono83/query/filters"
	"github.com/mono83/query/sorting"
)

// FilterLeft iterates over all rules (even nested) and replaces
// all string field names (left part of rule) with their query.Field
// definitions.
// Also, replace sorting settings.
// Method will fail on unknown fields.
func FilterLeft(f query.Filter, fs []query.Field) (query.Filter, error) {
	m := fieldNameMap(fs)
	boxed, err := ConditionLeft(f, fs)
	if err != nil {
		return nil, err
	}

	var srt []query.Sorting
	for _, s := range f.Sorting() {
		if x, ok := m[s.Name()]; ok {
			srt = append(srt, sorting.Field(x, s.Type()))
		} else {
			return nil, fmt.Errorf("unknown field %s in sorting", s.Name())
		}
	}
	return filters.FromCondition(boxed, srt, f.Limit(), f.Offset()), nil
}

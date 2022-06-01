package box

import (
	"github.com/mono83/query"
	"github.com/mono83/query/queries"
)

// QueryLeft iterates over all rules (even nested) and replaces
// all string field names (left part of rule) with their query.Field
// definitions.
// Also, replace sorting settings.
// Method will fail on unknown fields.
func QueryLeft(q query.Query, fs []query.Field) (query.Query, error) {
	boxed, err := FilterLeft(q, fs)
	if err != nil {
		return nil, err
	}
	return queries.FromFilter(boxed, q.Schema(), q.Columns()), nil
}

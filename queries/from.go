package queries

import "github.com/mono83/query"

// FromFilter constructs Query from given Filter.
func FromFilter(f query.Filter, schema query.Named, columns []query.Named) query.Query {
	if f == nil {
		return nil
	}

	return New(columns, schema, f.Type(), f.Rules(), f.Conditions(), f.Sorting(), f.Limit(), f.Offset())
}

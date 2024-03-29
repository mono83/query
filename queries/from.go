package queries

import "github.com/mono83/query"

// FromFilter constructs Query from given Filter.
func FromFilter(f query.Filter, schema query.Named, columns []query.Named) query.Query {
	if f == nil {
		return nil
	}

	return New(columns, schema, f.Type(), f.Rules(), f.Conditions(), f.Sorting(), f.Limit(), f.Offset())
}

// AllFromFilter constructs query from given filter that will return all columns.
func AllFromFilter(f query.Filter, schema query.Named) query.Query {
	return FromFilter(f, schema, nil)
}

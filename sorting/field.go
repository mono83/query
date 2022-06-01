package sorting

import "github.com/mono83/query"

// Field constructs sorting for given field
func Field(f query.Field, o query.SortOrder) query.Sorting {
	return field{Field: f, order: o}
}

type field struct {
	query.Field
	order query.SortOrder
}

func (f field) Type() query.SortOrder { return f.order }

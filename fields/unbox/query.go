package unbox

import (
	"github.com/mono83/query"
	"github.com/mono83/query/queries"
)

// Query iterates over sorting settings and rules (including
// nested) and replaces all found fields with their string
// name by calling field.Name method.
func Query(q query.Query) query.Query {
	return queries.FromFilter(Filter(q), q.Schema(), q.Columns())
}

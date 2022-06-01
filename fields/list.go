package fields

import "github.com/mono83/query"

// List is s handy func to create slice of fields.
func List(f ...query.Field) []query.Field { return f }

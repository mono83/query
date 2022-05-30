package query

// Aliased is interface for entries (columns, schemas) that can be aliased.
type Aliased interface {
	Alias() string
}

// Aliased defines pair of alias and original name
type AliasedName interface {
	Aliased
	Named
}

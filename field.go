package query

// Field defined physical or logical column
type Field interface {
	Named
	Sortable() bool
	Filterable() bool
}

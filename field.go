package query

// Field defined physical or logical column
type Field interface {
	Named
	DataType() interface{}
	Sortable() bool
	Filterable() bool
}

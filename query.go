package query

// Query describes common select query
type Query interface {
	Filter

	Schema() Named
	Columns() []Named
}

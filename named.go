package query

// Named is interface for named entries (columns, schemas)
type Named interface {
	Name() string
}

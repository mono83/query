package query

// Filter extends condition with limits and sorting
type Filter interface {
	Condition

	Sorting() []Sorting
	Limit() int
	Offset() int
}

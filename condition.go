package query

// Condition represents complex condition, that may include rules and other conditions
type Condition interface {
	Type() Logic
	Rules() []Rule
	Conditions() []Condition
}

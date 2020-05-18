package query

import "github.com/mono83/query/match"

// Rule contains rule definition.
// It consist of left and right parts and matching operator.
type Rule interface {
	GetLeft() interface{}
	GetRight() interface{}
	GetType() match.Type
}

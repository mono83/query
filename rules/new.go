package rules

import (
	"fmt"

	"github.com/mono83/query"
	"github.com/mono83/query/match"
)

// New builds new rule
func New(left interface{}, op match.Type, right interface{}) query.Rule {
	if right == nil {
		return leftPart{l: left, t: op}
	}
	return full{l: left, r: right, t: op}
}

type full struct {
	l, r interface{}
	t    match.Type
}

func (f full) GetLeft() interface{}  { return f.l }
func (f full) GetRight() interface{} { return f.r }
func (f full) GetType() match.Type   { return f.t }
func (f full) String() string {
	return fmt.Sprintf(`{Rule {%v (%T)} %s {%v (%T)}}`, f.l, f.l, f.t.String(), f.r, f.r)
}

type leftPart struct {
	l interface{}
	t match.Type
}

func (l leftPart) GetLeft() interface{} { return l.l }
func (leftPart) GetRight() interface{}  { return nil }
func (l leftPart) GetType() match.Type  { return l.t }
func (l leftPart) String() string {
	return fmt.Sprintf(`{Rule {%v (%T)} %s}}`, l.l, l.l, l.t.String())
}

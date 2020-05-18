package query

import (
	"fmt"

	"github.com/mono83/query/match"
)

// CommonRule is simple Rule implementation
type CommonRule struct {
	Left  interface{}
	Type  match.Type
	Right interface{}
}

// GetLeft returns left part of rule condition
func (c CommonRule) GetLeft() interface{} { return c.Left }

// GetRight returns right part of rule condition
func (c CommonRule) GetRight() interface{} { return c.Right }

// GetType return operation, used in CommonRule
func (c CommonRule) GetType() match.Type { return c.Type }

func (c CommonRule) String() string {
	return fmt.Sprintf(
		`{Rule {%v (%T)} %s {%v (%T)}}`,
		c.Left,
		c.Left,
		c.Type.String(),
		c.Right,
		c.Right,
	)
}

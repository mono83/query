package names

import "github.com/mono83/query"

// String creates query.Named implementation using plain string
func String(s string) query.Named {
	return named(s)
}

type named string

func (n named) Name() string   { return string(n) }
func (n named) String() string { return string(n) }

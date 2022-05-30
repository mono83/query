package names

import "github.com/mono83/query"

// Alias constructs query.Aliased instance
func Alias(name, alias string) query.AliasedName {
	return aliased{name, alias}
}

type aliased [2]string

func (a aliased) Name() string   { return a[0] }
func (a aliased) Alias() string  { return a[1] }
func (a aliased) String() string { return a[0] + " as " + a[1] }

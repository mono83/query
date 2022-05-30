package sorting

import "github.com/mono83/query"

// Asc constructs ascending sorting
func Asc(s string) query.Sorting { return String(s, query.Asc) }

// Desc constructs descending sorting
func Desc(s string) query.Sorting { return String(s, query.Desc) }

// String constructs new sorting qualifier with
// requested order.
func String(s string, o query.SortOrder) query.Sorting {
	if o == query.Asc {
		return sortingAsc(s)
	} else if o == query.Desc {
		return sortingDesc(s)
	}

	return sortingOther(s)
}

type sortingAsc string
type sortingDesc string
type sortingOther string

func (sortingAsc) Type() query.SortOrder   { return query.Asc }
func (sortingDesc) Type() query.SortOrder  { return query.Desc }
func (sortingOther) Type() query.SortOrder { return query.UnknownSort }

func (s sortingAsc) Name() string   { return string(s) }
func (s sortingDesc) Name() string  { return string(s) }
func (s sortingOther) Name() string { return string(s) }

func (s sortingAsc) String() string   { return string(s) }
func (s sortingDesc) String() string  { return string(s) }
func (s sortingOther) String() string { return string(s) }

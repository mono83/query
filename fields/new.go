package fields

import "github.com/mono83/query"

// New constructs new field definition with given settings
func New(name string, filter, sort bool, dataType interface{}) query.Field {
	return field{name: name, sort: sort, filter: filter, dataType: dataType}
}

// Untyped constructs new field definition with given settings but without type
func Untyped(name string, filter, sort bool) query.Field {
	return field{name: name, sort: sort, filter: filter}
}

// Typed adds type definition for given field
func Typed(f query.Field, dataType interface{}) query.Field {
	return New(f.Name(), f.Filterable(), f.Sortable(), dataType)
}

// Indexed constructs new field definition with enabled
// sorting and filtering capabilities.
func Indexed(name string) query.Field {
	return Untyped(name, true, true)
}

// Filterable constructs new field definition with only
// filtering capability. It can be suitable for enums.
func Filterable(name string) query.Field {
	return Untyped(name, true, false)
}

type field struct {
	dataType     interface{}
	name         string
	sort, filter bool
}

func (f field) DataType() interface{} { return f.dataType }
func (f field) Name() string          { return f.name }
func (f field) Sortable() bool        { return f.sort }
func (f field) Filterable() bool      { return f.filter }

func (f field) String() string {
	if f.sort && f.filter {
		return `{"` + f.name + `",filter,sort}`
	}
	if !f.sort && !f.filter {
		return `{"` + f.name + `",nofilter,nosort}`
	}
	if f.sort {
		return `{"` + f.name + `",nofilter,sort}`
	}
	return `{"` + f.name + `",filter,nosort}`
}

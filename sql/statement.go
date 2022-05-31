package sql

// Statement is an interface to SQL with placeholders.
type Statement interface {
	Query() string
	Args() []interface{}
}

// NewStatement constructs new statement for given query
// and arguments.
func NewStatement(query string, args ...interface{}) Statement {
	if len(args) == 0 {
		return stringStatement(query)
	}

	return statement{query: query, args: args}
}

type statement struct {
	query string
	args  []interface{}
}

func (s statement) Query() string       { return s.query }
func (s statement) Args() []interface{} { return s.args }

type stringStatement string

func (s stringStatement) Query() string       { return string(s) }
func (s stringStatement) Args() []interface{} { return nil }

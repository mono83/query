package mysql

import (
	"bytes"
	"errors"
	"strings"

	"github.com/mono83/query"
	"github.com/mono83/query/sql"
)

// StatementBuilder is component, used to create statements from conditions and filters
type StatementBuilder struct {
	buf          *bytes.Buffer
	placeholders []interface{}
}

// NewStatementBuilder returns new StatementBuilder struct
func NewStatementBuilder() *StatementBuilder {
	return &StatementBuilder{buf: bytes.NewBuffer(nil)}
}

// Build returns statement
func (s *StatementBuilder) Build() sql.Statement {
	return sql.NewStatement(s.buf.String(), s.placeholders...)
}

// WriteKey writes table or column name
func (s *StatementBuilder) WriteKey(key string) *StatementBuilder {
	if l := len(key); l > 2 {
		if key[0] == '`' && key[l-1] == '`' && strings.Count(key, "`") == 2 {
			s.buf.WriteString(key)
			return s
		}
	}

	s.buf.WriteRune('`')
	s.buf.WriteString(key)
	s.buf.WriteRune('`')
	return s
}

// WriteNamed writes named entity
func (s *StatementBuilder) WriteNamed(n query.Named) *StatementBuilder {
	if n != nil {
		s.WriteKey(n.Name())
	}

	return s
}

// WriteColumn writes column name, aliases are supported
func (s *StatementBuilder) WriteColumn(n query.Named) error {
	if n != nil {
		s.WriteKey(n.Name())

		if a, ok := n.(query.Aliased); ok {
			s.buf.WriteString(" as ")
			s.WriteKey(a.Alias())
		}
	} else {
		return errors.New("nil provided instead column")
	}

	return nil
}

// WriteSchema writes schema (table) name, aliases not supported
func (s *StatementBuilder) WriteSchema(n query.Named) error {
	if n != nil {
		s.WriteKey(n.Name())
	} else {
		return errors.New("nil provided instead schema")
	}

	return nil
}

package mysql

import (
	"errors"

	"github.com/mono83/query"
)

// WriteQuery writes whole query into statement builder
func (s *StatementBuilder) WriteQuery(q query.Query) error {
	if q == nil {
		return errors.New("nil provided instead Query")
	}

	// Writing SELECT
	s.buf.WriteString("SELECT ")

	// Writing columns
	if len(q.Columns()) == 0 {
		s.buf.WriteString("*")
	} else {
		for i, c := range q.Columns() {
			if i > 0 {
				s.buf.WriteString(", ")
			}
			if err := s.WriteColumn(c); err != nil {
				return err
			}
		}
	}

	// Writing FROM
	s.buf.WriteString(" FROM ")
	if err := s.WriteSchema(q.Schema()); err != nil {
		return err
	}

	// Writing filter
	if len(q.Conditions()) != 0 || len(q.Rules()) != 0 {
		s.buf.WriteString(" WHERE ")
	}
	return s.WriteFilter(q)
}

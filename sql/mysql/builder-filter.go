package mysql

import (
	"errors"
	"strconv"

	"github.com/mono83/query"
)

// WriteFilter converts filter into SQL and writes it into buffer
func (s *StatementBuilder) WriteFilter(f query.Filter) error {
	if err := s.WriteCondition(f); err != nil {
		return err
	}

	if len(f.GetSorting()) > 0 {
		// Applying sorting
		s.buf.WriteString(" ORDER BY ")

		for i, sort := range f.GetSorting() {
			if i > 0 {
				s.buf.WriteRune(',')
			}
			s.WriteNamed(sort)
			s.buf.WriteRune(' ')
			if sort.Type() == query.Desc {
				s.buf.WriteString("DESC")
			} else if sort.Type() == query.Asc {
				s.buf.WriteString("ASC")
			} else {
				return errors.New("unknown sort type")
			}
		}
	}

	if lim := f.GetLimit(); lim > 0 {
		// Applying limit
		s.buf.WriteString(" LIMIT ")
		if off := f.GetOffset(); off > 0 {
			s.buf.WriteString(strconv.Itoa(off))
			s.buf.WriteRune(',')
		}
		s.buf.WriteString(strconv.Itoa(lim))
	}

	return nil
}

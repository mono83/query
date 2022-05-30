package mysql

import (
	"errors"

	"github.com/mono83/query"
)

// WriteCondition converts condition into SQL and writes it into buffer
func (s *StatementBuilder) WriteCondition(cond query.Condition) error {
	if len(cond.Conditions()) == 0 && len(cond.Rules()) == 0 {
		return errors.New("empty condition - it has no rules and nested conditions")
	} else if len(cond.Conditions()) == 0 && len(cond.Rules()) == 1 {
		return s.WriteRule(cond.Rules()[0])
	} else if len(cond.Rules()) == 0 && len(cond.Conditions()) == 1 {
		return s.WriteCondition(cond.Conditions()[0])
	}

	sep := ""
	if cond.Type() == query.Or {
		sep = " OR "
	} else if cond.Type() == query.And {
		sep = " AND "
	} else {
		return errors.New("unsupported condition logic - it neither AND nor OR")
	}

	s.buf.WriteRune('(')
	i := 0

	for _, r := range cond.Rules() {
		if i > 0 {
			s.buf.WriteString(sep)
		}
		err := s.WriteRule(r)
		if err != nil {
			return err
		}
		i++
	}

	for _, c := range cond.Conditions() {
		if i > 0 {
			s.buf.WriteString(sep)
		}
		err := s.WriteCondition(c)
		if err != nil {
			return err
		}
		i++
	}

	s.buf.WriteRune(')')

	return nil
}

package goar

import (
	"fmt"
	"strings"
)

type Select struct {
	table   string
	columns []string
	where   *condition
}

func (s *Select) Table(table string) *Select {
	s.table = table
	return s
}

func (s *Select) Columns(columns []string) *Select {
	s.columns = columns
	return s
}

func (s *Select) Where(cond string, args ...interface{}) *Select {
	if s.where == nil {
		s.where = &condition{phrase: "WHERE"}
	}
	s.where.addExpression(cond, args)
	return s
}

func (s *Select) Build() (query string, binds []interface{}) {
	baseQuery := fmt.Sprintf("SELECT %s FROM %s", strings.Join(s.columns, ", "), s.table)
	whereQuery, whereBinds := s.where.build()
	return baseQuery + whereQuery, whereBinds
}

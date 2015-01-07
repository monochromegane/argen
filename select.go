package goar

import (
	"fmt"
	"strings"
)

type Select struct {
	table   string
	columns []string
	where   *Condition
}

func (s *Select) Table(table string) *Select {
	s.table = table
	return s
}

func (s *Select) Columns(columns []string) *Select {
	s.columns = columns
	return s
}

func (s *Select) Where(conditions []ConditionParam) *Select {
	if s.where == nil {
		s.where = &Condition{Phrase: "WHERE"}
	}
	s.where.SetConditions(conditions)
	return s
}

func (s *Select) Build() (query string, binds []interface{}) {
	baseQuery := fmt.Sprintf("SELECT %s FROM %s", strings.Join(s.columns, ", "), s.table)
	whereQuery, whereBinds := s.where.Build()
	return baseQuery + whereQuery, whereBinds
}

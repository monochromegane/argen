package goar

import (
	"fmt"
	"strings"
)

type Select struct {
	table   string
	columns []string
	orderBy *orderBy
	limit   *limit
	where   *condition
}

func (s *Select) Table(table string) *Select {
	s.table = table
	return s
}

func (s *Select) Columns(columns ...string) *Select {
	s.columns = columns
	return s
}

func (s *Select) Where(cond string, args ...interface{}) *Select {
	if s.where == nil {
		s.where = &condition{phrase: "WHERE"}
	}
	s.where.addExpression(cond, args...)
	return s
}

func (s *Select) And(cond string, args ...interface{}) *Select {
	return s.Where(cond, args...)
}

func (s *Select) OrderBy(column, order string) *Select {
	s.orderBy.addOrder(column, order)
	return s
}

func (s *Select) Limit(limit int) *Select {
	s.limit.setLimit(limit)
	return s
}

func (s *Select) Build() (query string, binds []interface{}) {
	baseQuery := fmt.Sprintf("SELECT %s FROM %s", strings.Join(s.columns, ", "), s.table)
	whereQuery, whereBinds := s.where.build()
	return baseQuery + whereQuery, whereBinds
}

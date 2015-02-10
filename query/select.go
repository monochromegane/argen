package query

import (
	"fmt"
	"strings"
)

type Select struct {
	table   string
	columns []string
	orderBy *orderBy
	limit   *limit
	offset  *offset
	groupBy *groupBy
	where   *condition
	having  *condition
	joins   []*join
	explain bool
}

func (s *Select) Table(table string) *Select {
	s.table = table
	return s
}

func (s *Select) Columns(columns ...string) *Select {
	s.columns = columns
	return s
}

func (s *Select) GetColumns() []string {
	return s.columns
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

func (s *Select) Offset(offset int) *Select {
	s.offset.setOffset(offset)
	return s
}

func (s *Select) GroupBy(group string, groups ...string) *Select {
	s.groupBy.setGroups(group, groups...)
	return s
}

func (s *Select) Having(cond string, args ...interface{}) *Select {
	if s.having == nil {
		s.having = &condition{phrase: "HAVING"}
	}
	s.having.addExpression(cond, args...)
	return s
}

func (s *Select) InnerJoin(table string, cond string, args ...interface{}) *Select {
	s.joins = append(s.joins, innerJoin(table, cond, args...))
	return s
}

func (s *Select) Explain() *Select {
	s.explain = true
	return s
}

func (s *Select) Build() (string, []interface{}) {
	baseQuery := fmt.Sprintf("SELECT %s FROM %s", strings.Join(s.columns, ", "), s.table)
	var binds []interface{}
	var explain string
	if s.explain {
		explain = "EXPLAIN "
	}
	var joinQuery string
	for _, j := range s.joins {
		q, b := j.build()
		joinQuery += q
		binds = append(binds, b...)
	}
	whereQuery, whereBinds := s.where.build()
	limitQuery, limitBinds := s.limit.build()
	offsetQuery, offsetBinds := s.offset.build()
	groupQuery := s.groupBy.build()
	havingQuery, havingBinds := s.having.build()
	orderQuery := s.orderBy.build()
	binds = append(binds, whereBinds...)
	binds = append(binds, limitBinds...)
	binds = append(binds, havingBinds...)
	binds = append(binds, offsetBinds...)
	return explain +
			baseQuery +
			joinQuery +
			whereQuery +
			limitQuery +
			offsetQuery +
			groupQuery +
			havingQuery +
			orderQuery,
		binds
}

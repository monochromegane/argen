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
	if s.orderBy == nil {
		s.orderBy = &orderBy{}
	}
	s.orderBy.addOrder(column, order)
	return s
}

func (s *Select) Limit(number int) *Select {
	if s.limit == nil {
		s.limit = &limit{}
	}
	s.limit.setLimit(number)
	return s
}

func (s *Select) Offset(number int) *Select {
	if s.offset == nil {
		s.offset = &offset{}
	}
	s.offset.setOffset(number)
	return s
}

func (s *Select) GroupBy(group string, groups ...string) *Select {
	if s.groupBy == nil {
		s.groupBy = &groupBy{}
	}
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

	var binds []interface{}

	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(s.columnsWithTable(), ", "), s.table)

	if s.explain {
		explain := "EXPLAIN "
		query = explain + query
	}

	var joinQuery string
	for _, j := range s.joins {
		q, b := j.build()
		joinQuery += q
		binds = append(binds, b...)
	}
	query += joinQuery

	if s.where != nil {
		q, b := s.where.build()
		query += q
		binds = append(binds, b...)
	}

	if s.groupBy != nil {
		q := s.groupBy.build()
		query += q
	}

	if s.having != nil {
		q, b := s.having.build()
		query += q
		binds = append(binds, b...)
	}

	if s.orderBy != nil {
		q := s.orderBy.build()
		query += q
	}

	if s.limit != nil {
		q, b := s.limit.build()
		query += q
		binds = append(binds, b...)
	}

	if s.offset != nil {
		if s.limit == nil {
			limit := limit{}
			limit.setLimit(-1)
			q, b := limit.build()
			query += q
			binds = append(binds, b...)
		}
		q, b := s.offset.build()
		query += q
		binds = append(binds, b...)
	}

	return query + ";", binds

}

func (s *Select) columnsWithTable() []string {
	columns := []string{}
	for _, c := range s.columns {
		if s.table == "" || strings.ContainsAny(c, ".") {
			columns = append(columns, c)
		} else {
			columns = append(columns, fmt.Sprintf("%s.%s", s.table, c))
		}
	}
	return columns
}

package query

import (
	"fmt"
	"strings"
)

type Update struct {
	table  string
	params Params
	where  *condition
}

func (u *Update) Table(table string) *Update {
	u.table = table
	return u
}

func (u *Update) Params(params Params) *Update {
	u.params = params
	return u
}

func (u *Update) Where(cond string, args ...interface{}) *Update {
	if u.where == nil {
		u.where = &condition{phrase: "WHERE"}
	}
	u.where.addExpression(cond, args...)
	return u
}

func (u *Update) And(cond string, args ...interface{}) *Update {
	return u.Where(cond, args...)
}

func (u *Update) Build() (string, []interface{}) {
	sets := []string{}
	binds := []interface{}{}

	for k, v := range u.params {
		sets = append(sets, fmt.Sprintf("%s = ?", k))
		binds = append(binds, v)
	}

	baseQuery := fmt.Sprintf("UPDATE %s SET %s", u.table, strings.Join(sets, ", "))

	whereQuery, whereBinds := u.where.build()
	return baseQuery + whereQuery, append(binds, whereBinds...)
}

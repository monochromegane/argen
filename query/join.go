package query

import "fmt"

type join struct {
	table string
	typ   string
	on    *condition
}

func innerJoin(table string, cond string, args ...interface{}) *join {
	j := &join{}
	return j.setJoin("INNER", table, cond, args...)
}

func (j *join) setJoin(typ string, table string, cond string, args ...interface{}) *join {
	j.typ = typ
	j.table = table
	if j.on == nil {
		j.on = &condition{phrase: "ON"}
	}
	j.on.addExpression(cond, args...)
	return j
}

func (j *join) build() (string, []interface{}) {
	baseQuery := fmt.Sprintf(" %s JOIN %s", j.typ, j.table)
	onQuery, onBinds := j.on.build()
	return baseQuery + onQuery, onBinds
}

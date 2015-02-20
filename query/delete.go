package query

import "fmt"

type Delete struct {
	table string
	where *condition
}

func (d *Delete) Table(table string) *Delete {
	d.table = table
	return d
}

func (d *Delete) Where(cond string, args ...interface{}) *Delete {
	if d.where == nil {
		d.where = &condition{phrase: "WHERE"}
	}
	d.where.addExpression(cond, args...)
	return d
}

func (d *Delete) And(cond string, args ...interface{}) *Delete {
	return d.Where(cond, args...)
}

func (d *Delete) Build() (string, []interface{}) {
	binds := []interface{}{}
	query := fmt.Sprintf("DELETE FROM %s", d.table)
	if d.where != nil {
		q, b := d.where.build()
		query += q
		binds = append(binds, b...)
	}
	return query + ";", binds
}

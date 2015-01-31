package gen

import "fmt"

type Templates []*Template

func (ts Templates) ToString() string {
	var template string
	for _, t := range ts {
		template = template + t.toDefine()
	}
	return template
}

type Template struct {
	Name string
	Text string
}

func (t Template) toDefine() string {
	return fmt.Sprintf("{{define \"%s\"}}%s{{end}}\n", t.Name, t.Text)
}

var templates = Templates{
	create,
	save,
	find,
	relation,
	query,
	queryRow,
	where,
	and,
	first,
	last,
	order,
	limit,
	offset,
	group,
	having,
	validation,
	hasMany,
	hasOne,
	belongsTo,
	scope,
}

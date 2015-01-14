package goar

import (
	"fmt"
	"strings"
)

type Insert struct {
	table  string
	params Params
}

type Params map[string]interface{}

func (i *Insert) Table(table string) *Insert {
	i.table = table
	return i
}

func (i *Insert) Params(params Params) *Insert {
	i.params = params
	return i
}

func (i *Insert) Build() (string, []interface{}) {
	columns := []string{}
	ph := []string{}
	binds := []interface{}{}

	for k, v := range i.params {
		columns = append(columns, k)
		ph = append(ph, "?")
		binds = append(binds, v)
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", i.table, strings.Join(columns, ", "), strings.Join(ph, ", ")), binds
}

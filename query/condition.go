package query

import (
	"fmt"
	"strings"
)

type condition struct {
	phrase      string
	expressions []expression
}

type expression struct {
	cond string
	args []interface{}
}

func (e expression) build() (string, []interface{}) {
	var query string
	var binds []interface{}
	switch len(e.args) {
	case 0:
		query = e.cond
	case 1:
		query = fmt.Sprintf("%s = ?", e.cond)
		binds = append(binds, e.args[0])
	case 2:
		query = fmt.Sprintf("%s %s ?", e.cond, e.args[0])
		binds = append(binds, e.args[1])
	default:
		query = ""
	}
	return query, binds
}

func (c *condition) addExpression(cond string, args ...interface{}) {
	c.expressions = append(c.expressions, expression{cond, args})
}

func (c *condition) build() (string, []interface{}) {
	var queries []string
	var binds []interface{}
	for _, e := range c.expressions {
		q, b := e.build()
		queries = append(queries, q)
		binds = append(binds, b...)
	}
	return fmt.Sprintf(" %s %s", c.phrase, strings.Join(queries, " AND ")), binds
}

package query

import (
	"fmt"
	"strings"
)

type condition struct {
	phrase  string
	queries []string
	binds   []interface{}
}

func (c *condition) addExpression(cond string, args ...interface{}) {
	var query string
	var binds []interface{}
	switch len(args) {
	case 0:
		query = cond
	case 1:
		query = fmt.Sprintf("%s = ?", cond)
		binds = append(binds, args[0])
	case 2:
		query = fmt.Sprintf("%s %s ?", cond, args[0])
		binds = append(binds, args[1])
	default:
		query = ""
	}
	c.queries = append(c.queries, query)
	c.binds = append(c.binds, binds...)
}

func (c *condition) build() (query string, binds []interface{}) {
	return fmt.Sprintf(" %s %s", c.phrase, strings.Join(c.queries, " AND ")), c.binds
}

package goar

import (
	"fmt"
	"strings"
)

type Condition struct {
	Phrase  string
	Queries []string
	Binds   []interface{}
}

type ConditionParam struct {
	Condition string
	Args      []interface{}
}

func (c *Condition) SetConditions(params []ConditionParam) {
	for _, param := range params {
		c.AddExpression(param)
	}
}

func (c *Condition) AddExpression(param ConditionParam) {
	var query string
	var binds []interface{}
	switch len(param.Args) {
	case 0:
		query = param.Condition
	case 1:
		query = fmt.Sprintf("%s = ?", param.Condition)
		binds = append(binds, param.Args[0])
	case 2:
		query = fmt.Sprintf("%s %s ?", param.Condition, param.Args[0])
		binds = append(binds, param.Args[1])
	default:
		query = ""
	}
	c.Queries = append(c.Queries, query)
	c.Binds = append(c.Binds, binds...)
}

func (c *Condition) Build() (query string, binds []interface{}) {
	return fmt.Sprintf(" %s %s", c.Phrase, strings.Join(c.Queries, " AND ")), c.Binds
}

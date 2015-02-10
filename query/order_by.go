package query

import (
	"fmt"
	"strings"
)

const (
	ASC  = "ASC"
	DESC = "DESC"
)

type orderBy struct {
	orders []order
}

type order struct {
	column string
	sort   string
}

func (o order) build() string {
	return fmt.Sprintf("%s %s", o.column, o.sort)
}

func (o *orderBy) addOrder(column, sort string) {
	o.orders = append(o.orders, order{column, sort})
}

func (o *orderBy) build() string {
	queries := []string{}
	for _, o := range o.orders {
		queries = append(queries, o.build())
	}
	return fmt.Sprintf(" ORDER BY %s", strings.Join(queries, ", "))
}

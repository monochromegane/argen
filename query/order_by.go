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
	queries []string
}

func (o *orderBy) addOrder(column, order string) {
	o.queries = append(o.queries, fmt.Sprintf("%s %s", column, order))
}

func (o *orderBy) build() string {
	return fmt.Sprintf("ORDER BY %s", strings.Join(o.queries, ", "))
}

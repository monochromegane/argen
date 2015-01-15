package ar

import (
	"fmt"
	"strings"
)

type groupBy struct {
	queries []string
}

func (g *groupBy) setGroups(group string, groups ...string) {
	g.queries = append(g.queries, group)
	if len(groups) > 0 {
		g.queries = append(g.queries, groups...)
	}
}

func (g *groupBy) build() string {
	return fmt.Sprintf(" GROUP BY %s", strings.Join(g.queries, ", "))
}

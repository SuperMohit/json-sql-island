package jsm

import (
	"fmt"
	"strings"
)

// groups by column
type groupBy struct {
	columns     []string
	nextClauses []Clause
}

func (g *groupBy) Next(n Clause) *groupBy {
	g.nextClauses = append(g.nextClauses, n)
	return g
}

func (g *groupBy) GroupBy(columns ...string) *groupBy {
	g.columns = append(g.columns, columns...)
	return g
}

// builds group by sql expression and calls the child clauses next in order
func (g groupBy) Build(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf(" GROUP BY %s", strings.Join(g.columns, ",")))
	for _, c := range g.nextClauses {
		c.Build(sb)
	}
}

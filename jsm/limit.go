package jsm

import (
	"fmt"
	"strings"
)

// used for pagination and limits
type limit struct {
	limit       int
	offset      int
	nextClauses []Clause
}

func (l *limit) Next(n Clause) *limit {
	l.nextClauses = append(l.nextClauses, n)
	return l
}

func (l *limit) Limit(li, o int) *limit {
	l.limit = li
	l.offset = o
	return l
}

// builds limit and offset expression and calls the child clauses next in order
func (l limit) Build(sb *strings.Builder) {
	if l.limit > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT %d", l.limit))
	}
	if l.offset > 0 {
		sb.WriteString(fmt.Sprintf(" OFFSET %d", (l.offset-1)*l.limit))
	}

	for _, c := range l.nextClauses {
		c.Build(sb)
	}
}

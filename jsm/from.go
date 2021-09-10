package jsm

import (
	"fmt"
	"strings"
)

type QueryBuilder interface {
	Build(sb strings.Builder)
}

// form clause takes list of tables, this can be extended to include select clause
type from struct {
	tables      []string
	nextClauses []Clause
}

func (f *from) From(next interface{}) *from {
	switch next.(type) {
	case []string:
		f.tables = next.([]string)
		return f
	}
	return f
}

func (f *from) Next(n Clause) *from {
	f.nextClauses = append(f.nextClauses, n)
	return f
}

// builds from sql expression and calls the child clauses next in order
func (f from) Build(sb *strings.Builder) {
	if len(f.tables) > 0 {
		sb.WriteString(fmt.Sprintf("FROM %s", strings.Join(f.tables, ", ")))
	}

	for _, c := range f.nextClauses {
		c.Build(sb)
	}
}

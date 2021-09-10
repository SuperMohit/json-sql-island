package jsm

import (
	"fmt"
	"strings"
)

// selects the column from the table, also can take care of count
type selectt struct {
	columns []string
	// from    *from
	typp        int
	nextClauses []Clause
}

type Clause interface {
	Build(sb *strings.Builder)
}

const (
	typeSelect = iota
	typeCount
)

func (s *selectt) Next(n Clause) *selectt {
	s.nextClauses = append(s.nextClauses, n)
	return s
}

func (s *selectt) Select(se interface{}) *selectt {
	switch se.(type) {
	case []string:
		sel := se.([]string)
		s.columns = append(s.columns, sel...)
		return s
	}
	return s
}

// builds select expression and calls the child clauses next in order
func (s *selectt) Build(sb *strings.Builder) {
	switch s.typp {
	case typeSelect:
		sb.WriteString(fmt.Sprintf("SELECT %s ", strings.Join(s.columns, ", ")))
	case typeCount:
		// TODO extend for specific column instead of *
		sb.WriteString(fmt.Sprintf("SELECT COUNT(*) "))
	}

	// Build expression tree for next clauses
	for _, c := range s.nextClauses {
		c.Build(sb)
	}
}

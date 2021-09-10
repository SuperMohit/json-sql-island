package jsm

import (
	"fmt"
	"strings"
)

type operator string

const (
	conditionEquals operator = "="
	conditionAfter  operator = ">="
	conditionBefore operator = "<="
	conditionIN     operator = "IN"

	// TODO implementation
	conditionBetween operator = "BETWEEN"
)

// where takes in the conditions to filter data from the table
type where struct {
	conditions  []condition
	nextClauses []Clause
}

type condition struct {
	col   string
	op    string
	value interface{}
}

func (w *where) Where(key, value string, op operator) *where {
	w.conditions = append(w.conditions, condition{
		col:   key,
		op:    string(op),
		value: value,
	})
	return w
}

// can be used later for chaining
func (w *where) Next(n Clause) *where {
	w.nextClauses = append(w.nextClauses, n)
	return w
}

// builds where expression and calls the child clauses next in order
func (w where) Build(sb *strings.Builder) {
	sb.WriteString(" WHERE")
	l := len(w.conditions) - 1
	for i, w := range w.conditions {
		if w.op == string(conditionIN) {
			// TODO based on datatype
			sb.WriteString(fmt.Sprintf(" %s %s (%s)", w.col, w.op, w.value))
		} else {
			sb.WriteString(fmt.Sprintf(" %s %s %s", w.col, w.op, w.value))
		}
		if i < l {
			// TODO for OR etc
			sb.WriteString(" AND")
		}
	}

	for _, c := range w.nextClauses {
		c.Build(sb)
	}
}

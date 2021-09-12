package jsm

import (
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
	operatorMap map[string]func(column string, values interface{}) condition
	// initialize for specific function
	sqlBuilder func(sb *strings.Builder)
}

type condition struct {
	col   string
	op    string
	value interface{}
}

func (w *where) Where(key string, values interface{}, op string) *where {

	if op != string(conditionBetween) {
		w.conditions = append(w.conditions, condition{
			col:   key,
			op:    op,
			value: values,
		})
		return w
	}

	transformer := w.operatorMap[op]
	c := transformer(key, values)
	w.conditions = append(w.conditions, c)
	return w
}

// can be used later for chaining
func (w *where) Next(n Clause) *where {
	w.nextClauses = append(w.nextClauses, n)
	return w
}

// builds where expression and calls the child clauses next in order
func (w where) Build(sb *strings.Builder) {
	w.sqlBuilder(sb)
	for _, c := range w.nextClauses {
		c.Build(sb)
	}
}

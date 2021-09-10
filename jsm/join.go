package jsm

import (
	"fmt"
	"strings"
)

type joinType string

const (
	LeftJOIN  joinType = "LEFT"
	RightJOIN joinType = "RIGHT"
	InnerJOIN joinType = "INNER"
	CrossJOIN joinType = "CROSS"
)

// join tales jointype table names and join conditions,
// this can also be extended to include the select clause instead of table
type join struct {
	typ         joinType
	table       string
	on          string
	nextClauses []Clause
}

type joins struct {
	joinClauses []join
	nextClauses []Clause
}

func mapJoinType(j string) joinType {
	switch j {
	case "LEFT":
		return LeftJOIN
	case "RIGHT":
		return RightJOIN
	case "INNER":
		return InnerJOIN
	case "CROSS":
		return CrossJOIN
	}
	return ""
}

func (j *joins) Join(typ joinType, table, on string) *joins {
	j.joinClauses = append(j.joinClauses, join{
		typ:   typ,
		table: table,
		on:    on,
	})
	return j
}

func (j *joins) Next(n Clause) *joins {
	j.nextClauses = append(j.nextClauses, n)
	return j
}

// builds join sql expression and calls the child clauses next in order
func (j *joins) Build(sb *strings.Builder) {
	for _, v := range j.joinClauses {
		sb.WriteString(fmt.Sprintf(" %s JOIN %s ON %s", v.typ, v.table, v.on))
	}

	for _, c := range j.nextClauses {
		c.Build(sb)
	}
}

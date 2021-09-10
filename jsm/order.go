package jsm

import (
	"strings"
)

type columnOrder struct {
	column string
	desc   bool
}

// orders the result set with the specified column and either ASC or DESC
type order struct {
	columnOrders []columnOrder
}

// builds order by expression and calls the child clauses next in order
func (o order) Build(sb *strings.Builder) {
	sb.WriteString(" ORDER BY ")
	l := len(o.columnOrders) - 1
	for i, o := range o.columnOrders {
		sb.WriteString(o.column)
		if o.desc {
			sb.WriteString(" DESC")
		}

		if i < l {
			sb.WriteString(",")
		}
	}
}

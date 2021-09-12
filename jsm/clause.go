package jsm

import (
	"fmt"
	"strings"
)

type dataBase int

const (
	postgres dataBase = iota
	mysql             = iota
)

// Factory for building specific nodes for the parse tree
type clause struct {
	*where
	// lets assume the where expression is specific to DB
	sbSpecificWhereSql map[dataBase]func(sb *strings.Builder)
}

func NewClause() *clause {
	mWhere := DbSpecificQuery()
	w := where{
		operatorMap: func() map[string]func(column string, values interface{}) condition {
			m := make(map[string]func(column string, values interface{}) condition)
			m[string(conditionBetween)] = func(column string, values interface{}) condition {
				return condition{
					col: column,
					op:  string(conditionBetween),
					//type validation for arrays
					value: func(values interface{}) interface{} {
						switch values.(type) {
						case []interface{}:
							ts := values.([]interface{})
							vs := make([]float64, 0, len(ts))
							for _, i := range ts {
								vs = append(vs, i.(float64))
							}
							return fmt.Sprintf("%v AND %v", vs[0], vs[1])
						}
						return nil
					}(values),
				}
			}

			return m
		}(),
	}

	w.sqlBuilder = func(sb *strings.Builder) {
		//db specific thing can be stored in map
		mWhere[postgres](sb)
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

	}
	return &clause{&w, mWhere}
}

// function that transforms the expressions based on the query
func DbSpecificQuery() map[dataBase]func(sb *strings.Builder) {
	mWhere := make(map[dataBase]func(sb *strings.Builder))
	mWhere[postgres] = func(sb *strings.Builder) {
		sb.WriteString(" WHERE")
	}

	mWhere[mysql] = func(sb *strings.Builder) {
		sb.WriteString(" WHERE")
	}
	return mWhere
}

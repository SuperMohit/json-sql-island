package jsm

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/SuperMohit/json-sql-island/jsm/resources"
	"reflect"
	"strings"
)

type QueryParser struct {
	query     resources.QuerySchema
	parseTree Clause
	sql       string
}

//go:embed resources/input.json
var query []byte

func (q *QueryParser) Parse() (string, error) {
	err := q.jsonReader()
	if err != nil {
		return "", err
	}
	err = q.generateParseTree()
	if err != nil {
		return "", err
	}
	err = q.buildExpression()
	if err != nil {
		return "", err
	}
	return q.sql, nil
}

func (q *QueryParser) jsonReader() error {
	err := json.Unmarshal(query, &q.query)
	if err != nil {
		return fmt.Errorf("error parsing the input sql %w", err)
	}
	return nil
}

func (q *QueryParser) buildExpression() error {
	sb := strings.Builder{}
	q.parseTree.Build(&sb)
	q.sql = sb.String()
	return nil
}

func (q *QueryParser) generateParseTree() error {
	s1 := reflect.ValueOf(&q.query).Elem()
	typeOfT := s1.Type()
	var nextClause Clause
	for i := s1.NumField() - 1; i >= 0; i-- {
		f := s1.Field(i)
		name := typeOfT.Field(i).Name
		switch name {
		case "Orderby":
			if !f.IsZero() {
				o := order{
					columnOrders: func() []columnOrder {
						co := make([]columnOrder, 0, len(q.query.Orderby.Columns))
						for _, v := range q.query.Orderby.Columns {
							co = append(co, columnOrder{
								column: v.Name,
								desc:   v.Desc,
							})
						}
						return co
					}(),
				}
				nextClause = o
			}
		case "Join":
			if !f.IsZero() {
				var jos []join
				for _, j := range q.query.Join {
					jo := join{
						table: j.Table,
						on:    j.On,
						// nextClauses: q.nextClause(nextClause)(),
					}
					jo.typ = mapJoinType(j.Type)
					jos = append(jos, jo)
				}
				nextClause = &joins{
					joinClauses: jos,
					nextClauses: q.nextClause(nextClause)(),
				}
			}
		case "Where":
			if !f.IsZero() {
				w := where{
					conditions: func() []condition {
						conds := make([]condition, 0, len(q.query.Where))
						for _, v := range q.query.Where {
							conds = append(conds,
								condition{
									col:   v.Fieldname,
									op:    v.Operator,
									value: v.Fieldvalue,
								})
						}
						return conds
					}(),
					nextClauses: q.nextClause(nextClause)(),
				}
				nextClause = w
			}
		case "From":
			if !f.IsZero() {
				f := from{
					tables:      q.query.From.Tables,
					nextClauses: q.nextClause(nextClause)(),
				}
				nextClause = f
			}
		case "Select":
			if !f.IsZero() {
				s := selectt{
					columns:     q.query.Select.Columns,
					nextClauses: q.nextClause(nextClause)(),
				}
				q.parseTree = &s
				break
			}
		case "Group":
			{
				if !f.IsZero() {
					g := groupBy{
						columns:     q.query.Group,
						nextClauses: q.nextClause(nextClause)(),
					}
					nextClause = g
				}
			}
		}
	}

	return nil
}

func (q *QueryParser) nextClause(nextClause Clause) func() []Clause {
	return func() []Clause {
		if nextClause != nil {
			return []Clause{nextClause}
		}
		return nil
	}
}

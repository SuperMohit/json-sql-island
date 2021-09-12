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
	query       resources.QuerySchema
	parseTree   Clause
	sql         string
	isSubClause bool
	*clause
}

func NewQueryParser(cl *clause) *QueryParser {
	return &QueryParser{
		clause: cl,
	}
}

//go:embed resources/input.json
var query []byte

func (q *QueryParser) Parse(body []byte) (string, error) {
	// default input for demo
	if len(body) == 0 {
		body = query
	}

	// Reading the json
	err := q.jsonReader(body)
	if err != nil {
		return "", err
	}

	// generate the parse tree
	err = q.generateParseTree()
	if err != nil {
		return "", err
	}

	//Build the expression
	err = q.buildExpression()
	if err != nil {
		return "", err
	}
	return q.sql, nil
}

// not implemented validator=-- schema validation
func (q *QueryParser) jsonReader(body []byte) error {
	err := json.Unmarshal(body, &q.query)
	if err != nil {
		return fmt.Errorf("error parsing the input sql %w", err)
	}
	return nil
}

func (q *QueryParser) buildExpression() error {
	sb := strings.Builder{}
	if q.isSubClause {
		sb.WriteString("(")
	}
	q.parseTree.Build(&sb)
	if q.isSubClause {
		sb.WriteString(")")
	}
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
				w := q.clause.where
				w.conditions = func() []condition {
					for _, v := range q.query.Where {
						value := v.Fieldvalue
						switch value.(type) {
						case string, int, []interface{}:
							// other scalar data types
							break
						default:
							value = recurseSubClauses(value)
						}
						w.Where(v.Fieldname, value, v.Operator)

					}
					return w.conditions
				}()

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

func recurseSubClauses(value interface{}) interface{} {
	jsonStr, _ := json.Marshal(value)

	cl := NewClause()
	parser := NewQueryParser(cl)
	value, err := parser.Parse(jsonStr)
	if err != nil {
		fmt.Errorf("error marshalling data %v", err)
	}
	return value
}

func (q *QueryParser) nextClause(nextClause Clause) func() []Clause {
	return func() []Clause {
		if nextClause != nil {
			return []Clause{nextClause}
		}
		return nil
	}
}

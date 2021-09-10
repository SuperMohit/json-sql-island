package main

import (
	"fmt"
	"github.com/SuperMohit/json-sql-island/jsm"
)

// This is a sample implementation of a SQL parser from a JSON file
// This uses Interpreter pattern
// First, it would build a parse or syntax Tree.
// Second, it would traverse the syntax tree and build the expression for the SQL
// Print the SQL to the console
func main() {
	parser := jsm.QueryParser{}
	q, _ := parser.Parse()
	fmt.Print(q)
}

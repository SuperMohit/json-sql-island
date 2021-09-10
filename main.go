package main

import (
	"flag"
	"fmt"
	"github.com/SuperMohit/json-sql-island/jsm"
	"io/ioutil"
)

// This is a sample implementation of a SQL parser from a JSON file
// This uses Interpreter pattern
// First, it would build a parse or syntax Tree.
// Second, it would traverse the syntax tree and build the expression for the SQL
// Print the SQL to the console
func main() {

	fptr := flag.String("fpath", "input.json", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	parser := jsm.QueryParser{}
	q, _ := parser.Parse(data)
	fmt.Print(q)
}

# Sql query from json input

* This is a sample implementation of a SQL parser from a JSON file
* This uses Interpreter pattern
* First, it would build a parse or syntax Tree.
* Second, it would traverse the syntax tree and build the expression for the SQL
* Print the SQL to the console

json file is stored at 
```
jsm/resources/input.json

```

# Execute from terminal
```
 % go run github.com/SuperMohit/json-sql-island 
SELECT column1, column2 FROM cayman WHERE column1 IN ('value1', 'value2', 'value3') AND column2 = value LEFT JOIN island ON cayman.id = island.id GROUP BY column1, column2 ORDER BY column1 DESC%  

```



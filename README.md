# Sql query from json input

* This is a sample implementation of a SQL parser from a JSON file
* Application is built using Interpreter pattern. 
* Each clause is a node of the Parse tree and the child nodes are next clauses. For instance SELECT is a parent node which would then for the child node like FROM  
* First, it would build a parse or syntax Tree.
* Second, it would traverse the syntax tree and build the expression for the SQL
* Print the SQL to the console

Input json file can be supplied as input to the command line. Default is input.json. 
There is a default json stored in resources which would be executed incase no input.json is supplied to command line.
```
jsm/resources/input.json

```

# Execute from terminal
```
 % go run github.com/SuperMohit/json-sql-island 
SELECT column1, column2 FROM cayman WHERE column1 IN ('value1', 'value2', 'value3') AND column2 = value LEFT JOIN island ON cayman.id = island.id GROUP BY column1, column2 ORDER BY column1 DESC%  

```
# godoc is generated at docs folder

![alt text](https://github.com/SuperMohit/json-sql-island/blob/master/doc.png)



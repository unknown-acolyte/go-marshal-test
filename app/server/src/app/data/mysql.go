package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// todo set pool options
//func main() {
var username = os.Getenv("marshalling_username")
var password = os.Getenv("marshalling_username")
var schema = os.Getenv("marshalling_schema")
var connectionString = fmt.Sprintf("%s:%s@tcp/%s", username, password, schema)
var DB, _ = sql.Open("mysql", connectionString)

//	if err != nil {
//		panic(error.Error)
//	}
//
//	defer Connector.Close() // what is this deferring until method exits
//
//	err = Connector.Ping()
//	if err != nil {
//		panic(err.Error)
//	}
//}

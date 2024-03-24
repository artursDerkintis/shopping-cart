package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:****@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Successfully connected to MySQL database")

}

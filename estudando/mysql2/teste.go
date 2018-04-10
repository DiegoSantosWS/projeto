package main

import (
	"database/sql"
	"fmt"
	/**
	github.com/go-sql-driver/mysql not is used in apllication directamente
	*/
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/mysql")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		panic(err)
	}

	var result []string

	for rows.Next() {
		var databaseName string

		err = rows.Scan(&databaseName)
		if err != nil {
			panic(err)
		}

		result = append(result, databaseName)
		fmt.Println("\n", result)
	}

	//return result
}

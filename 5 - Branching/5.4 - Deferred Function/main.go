package main

import (
	"database/sql"
	"fmt"
)

func main() {
	// Example 1
	fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
	defer fmt.Println("defer 2")

	db, _ := sql.Open("driverName", "connection string")
	defer db.Close() // wil close last

	rows, _ := db.Query("some sql query here")
	defer rows.Close() // will close first
}

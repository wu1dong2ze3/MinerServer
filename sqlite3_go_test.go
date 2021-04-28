package main

import (
	"database/sql"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"os"
)

type age int

/**
一个多态的例子
*/
func main() {

	sqlite3.Version()
	//creat database
	db, err := sql.Open("sqlite3", "./MyDatabase.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open database: %v\n", err)
		os.Exit(1)
	}

	//creat table
	sql := "CREATE TABLE STUDENTS(ID INT PRIMARY KEY     NOT NULL, NAME           TEXT    NOT NULL,AGE            INT     NOT NULL);"
	_, err = db.Exec(sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creat TableBase: %v\n", err)
		os.Exit(1)
	}

	//insert
	sql = "INSERT INTO STUDENTS (ID,NAME,AGE) VALUES (1, 'LiLi', 15);INSERT INTO STUDENTS (ID,NAME,AGE) VALUES (2, 'Lucy', 16); "
	_, err = db.Exec(sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Insert: %v\n", err)
		os.Exit(1)
	}

	//select
	rows, err := db.Query("SELECT * FROM STUDENTS")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Select: %v\n", err)
		os.Exit(1)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Select Rows: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(id, name, age)
	}

	//delete
	_, err = db.Exec("DELETE FROM STUDENTS WHERE ID=1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Delete: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("End")

}

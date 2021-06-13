package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user = "root"
	db_pwd  = "A1sujith."
	db_addr = "localhost"
	db_db   = "mysqldb"
)

func main() {
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_pwd, db_addr, db_db)
	db, err := sql.Open("mysql", s)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//insertData(db)
	//displayData(db)
	//deleteData(1, db)
	updateData(2, "Vizag", db)
}
func updateData(i int, s string, db *sql.DB) {
	q := "update student set city=? where id = ?"
	del, err := db.Prepare(q)
	if err != nil {
		fmt.Println("some error")
	}
	defer del.Close()
	_, err2 := del.Exec(s, i)
	if err2 != nil {
		fmt.Println("some error2")
	} else {
		fmt.Println("Data Updated")
	}
}
func deleteData(i int, db *sql.DB) {
	q := "delete from student where id = ?"
	del, err := db.Prepare(q)
	if err != nil {
		fmt.Println("some error")
	}
	defer del.Close()
	_, err2 := del.Exec(1)
	if err2 != nil {
		fmt.Println("some error2")
	} else {
		fmt.Println("Data Deleted")
	}
}
func insertData(db *sql.DB) {
	q := "Insert into student(name,city) values(?,?)"
	insert, err := db.Prepare(q)
	if err != nil {
		fmt.Println("some error")
	}
	defer insert.Close()
	_, err2 := insert.Exec("varma", "Salur")
	if err2 != nil {
		fmt.Println("some error2")
	} else {
		fmt.Println("Data Inserted")
	}
}

func displayData(db *sql.DB) {
	l := "select * from student"
	res, err1 := db.Query(l)
	if err1 != nil {
		fmt.Println("some error4")
	}
	defer res.Close()

	for res.Next() {
		var i int
		var s1 string
		var s2 string
		err3 := res.Scan(&i, &s1, &s2)
		fmt.Println(i, s1, s2)

		if err3 != nil {
			fmt.Println("some error3")
		}
	}
	//fmt.Println(i, s1, s2)
}

package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:redhat@tcp(127.0.0.1:3306)/ivan?charset=utf8")
	check(err)

	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/book", book)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Successfully Connected!")
	check(err)
}

func book(w http.ResponseWriter, req *http.Request) {
	q := "SELECT id, name from book;"
	rows, err := db.Query(q)
	if err != nil {
		http.Error(w, "Table Not Found", http.StatusNotFound)
		return
	}

	defer rows.Close()

	var s, id, name string
	s = "RETRIEVED RECORDS:\n"

	// Query the rows
	for rows.Next() {
		err = rows.Scan(&id, &name)
		check(err)
		s += id + " " + name + "\n"
	}

	_, err = fmt.Fprintln(w, s)
	check(err)
}

func create(w http.ResponseWriter, req *http.Request) {
	q := "CREATE TABLE if not exists customer (name VARCHAR(20));"

	stmt, err := db.Prepare(q)
	check(err)

	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Customer Table Created!", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	q := "INSERT INTO customer VALUES (\"James\");"
	stmt, err := db.Prepare(q)
	check(err)

	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	q := "SELECT * from customer;"
	rows, err := db.Query(q)
	check(err)

	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	q := "UPDATE customer SET name=\"Jimmy\" WHERE name=\"James\";"
	stmt, err := db.Prepare(q)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	q := "DELETE FROM customer WHERE name=\"Jimmy\";"
	stmt, err := db.Prepare(q)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	q := "DROP TABLE customer;"
	stmt, err := db.Prepare(q)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

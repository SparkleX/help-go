package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root:1234@tcp(1.1.1.101:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping failed")
	}
	var (
		user string
		host string
	)
	rows, err := db.Query("select User, Host from user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user, &host)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user, host)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

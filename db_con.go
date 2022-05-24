package main

import (
	"database/sql"
	"fmt"
)

func getPeopleFromDB() (people []person) {
	db, err := sql.Open("sqlite3", "./system.db")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT ID,Name,Phone,Dob FROM People")
	if err != nil {
		panic(err)

	}
	for rows.Next() {
		var per person
		rows.Scan(&per.ID, &per.Name, &per.Phone, &per.Dob)
		fmt.Print("person: ?", per)
		people = append(people, per)
	}
	return people
}

package main

import (
	"database/sql"
	"fmt"
)

///establishing a database connection
func estDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./system.db")
	if err != nil {
		panic(err)
	}
	return db
}

///adding a new person to the database C
func addPersonToDB(per person) (success bool) {
	db := estDB()
	var st = "INSERT INTO People(Name,Phone,Dob) VALUES(?,?,?)"
	stmt, err := db.Prepare(st)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(per.Name, per.Phone, per.Dob)
	return (err == nil)
}

////getting all people from the database - R
func getPeopleFromDB() (people []person) {
	db := estDB()
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

func removePersonFromDB(id string) (success bool) {
	db := estDB()
	var st = "DELETE FROM People WHERE ID = ?"
	stmt, err := db.Prepare(st)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id)
	return (err == nil)
}

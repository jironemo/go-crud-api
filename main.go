package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type person struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Dob   string `json:"dob"`
	Phone string `json:"phone"`
}

var people = []person{}

func getPeople(c *gin.Context) {
	c.IndentedJSON(200, people)
}
func addPeople(c *gin.Context) {
	var newPerson person
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	people = append(people, newPerson)

	c.IndentedJSON(http.StatusCreated, newPerson)
}

func PassQuery(id string) (name string) {
	db, err := sql.Open("sqlite3", "./system.db")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT Name FROM People WHERE id=?", id)
	if err != nil {
		panic(err)

	}
	for rows.Next() {
		rows.Scan(&name)
	}
	return name
}

func main() {
	/**
	router := gin.Default()
	router.GET("/people", getPeople)
	router.POST("/people", addPeople)
	router.Run("localhost:3000")
	**/
	var g string = PassQuery("1")
	fmt.Print(g)
}

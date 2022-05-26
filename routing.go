package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Dob   string `json:"dob"`
	Phone string `json:"phone"`
}

////initialize the api
func startRouter() {
	router := gin.Default()
	router.GET("/people", getPeople)
	router.POST("/people/add", addPeople)
	router.DELETE("/people/remove/:id", removePerson)
	router.Run("localhost:3000")
}

/**
Get the list of people from the server using GET method
**/
func getPeople(c *gin.Context) {
	c.IndentedJSON(200, getPeopleFromDB())
}

/**
////Add people into the server/ database using the POST method
**/
func addPeople(c *gin.Context) {
	var newPerson = person{"3", "Johnathan", "1999-20-12", "09781787873"}
	fmt.Print(addPersonToDB(newPerson))
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newPerson)
}

////Remove people from the server using the ID
func removePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	removePersonFromDB(id)
	c.IndentedJSON(204, getPeopleFromDB())
}

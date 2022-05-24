package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var people []person

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
	router.POST("/people", addPeople)
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
	var newPerson person
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	people = append(people, newPerson)

	c.IndentedJSON(http.StatusCreated, newPerson)
}

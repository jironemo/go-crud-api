package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
}
func TestCreate(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := SetUpRouter()
	r.POST("/people/add", addPeople)
	newPerson := person{
		Name:  "John",
		Phone: "123456789",
		Dob:   "01/01/1990",
	}
	jsonValue, _ := json.Marshal(newPerson)
	req, _ := http.NewRequest("POST", "/people/add", bytes.NewBuffer((jsonValue)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func TestGet(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := SetUpRouter()
	r.GET("/people", getPeople)
	req, _ := http.NewRequest("GET", "/people", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var people []person
	json.Unmarshal(w.Body.Bytes(), &people)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, people)

}

func TestDelete(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := SetUpRouter()
	r.DELETE("/people/remove/:id", removePerson)
	req, _ := http.NewRequest("DELETE", "/people/remove/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

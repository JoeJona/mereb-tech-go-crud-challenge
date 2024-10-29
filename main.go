package main

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Hobbies []string `json"hobbies"`
}

var persons = []person {}

func getAllPersons(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, persons)
}

func getPersonByID(id string) (*person, error) {
	for i, t := range persons {
		if t.Id == id {
			fmt.Printf(t.Id)
			return &persons[i], nil
		}
	}

	return nil, errors.New("person not found")

}

func getPerson(context *gin.Context) {
	id := context.Param("id")
	person, err := getPersonByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Person Found By this Id"})
	}

	context.IndentedJSON(http.StatusOK, person)
}

func addPerson(context *gin.Context) {

	var newPerson person

	if err := context.BindJSON(&newPerson); err != nil {
		return 
	}

	newPerson.Id = uuid.NewString()

	persons = append(persons, newPerson)

	context.IndentedJSON(http.StatusCreated, newPerson)

}

func editPerson(context *gin.Context) {

	var editPerson person
	
	id := context.Param("id")
	person, err := getPersonByID(id)
	

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Person Found By this Id"})
	}

	if err := context.BindJSON(&editPerson); err != nil {
		return 
	}

	editPerson.Id = id

	*person = editPerson

	context.IndentedJSON(http.StatusOK, editPerson)

}

func deletePerson(context *gin.Context) {
	
	var id int
	personId := context.Param("id")
	person, err := getPersonByID(personId)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Person Found By this Id"})
	}

	for i, t := range persons {
		if t.Id == person.Id {
			id = i
		}
	}

	persons = append(persons[:id], persons[id+1:]...)

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})

}

func main() {

	router := gin.Default()

	router.GET("/person/all", getAllPersons)
	router.GET("/person/:id", getPerson)
	router.POST("/person", addPerson)
	router.PUT("/person/:id", editPerson)
	router.DELETE("/person/:id", deletePerson)

	router.Run("localhost:8080")

}

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (person Person) to_String() string {
	return string("Id: " + strconv.Itoa(person.Id) + "\nName: " + person.Name +
		"\nAge: " + strconv.Itoa(person.Age) + "\nAddress: " + person.Address +
		"\nWork: " + person.Work)
}

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}

type personList struct {
	Persons []Person `json:"persons"`
}

func (i *Person) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("Name is a required field")
	}
	return nil
}
func (*personList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Person) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         string
	BirthDate   time.Timer
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) pringFirstName() string {
	return m.FirstName
}

func main() {

	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	// log.Println("myVar is set to", myVar.FirstName)
	// log.Println("myVar2 is set to", myVar2.FirstName)
	log.Println("myVar is set to", myVar.pringFirstName())
	log.Println("myVar2 is set to", myVar2.pringFirstName())

	user := User{
		FirstName: "Trevor",
		LastName:  "Hiroshi",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}

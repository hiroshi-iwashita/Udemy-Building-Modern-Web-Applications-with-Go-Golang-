package main

import (
	"log"
)

func main() {
	var myString string
	myString = "Green"
	log.Println("mystring is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after call func mystring is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}

package main

import "fmt"

var myName string

func main() {
	fmt.Println("Hello world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye cruel world"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, "fsuhjsdf", theOtherThingThatWasSaid)

	setMyName("Hiroshi")
	name := getMyName()
	fmt.Println(name)

}

func saySomething() (string, string) {
	return "something", "else"
}

func setMyName(nameToSet string) {
	myName = nameToSet
}

func getMyName() string {
	return myName
}

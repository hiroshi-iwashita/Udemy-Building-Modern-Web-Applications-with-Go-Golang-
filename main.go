package main

// $ go mod init github.com/hiroshi-iwashita/udemy-buildingModernWebApplicationsWithGo

import (
	"log"

	"github.com/hiroshi-iwashita/udemy-buildingModernWebApplicationsWithGo/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}

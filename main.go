package main

import "log"

func main() {

	myVar := "horse"
	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is set to something else")
	}

	myNum := 100
	isTrue := false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	// var isTrue bool
	// isTrue = false

	// if isTrue == true {
	// 	log.Println("isTrue is", isTrue)
	// } else {
	// 	log.Println("isTrue is", isTrue)
	// }
}

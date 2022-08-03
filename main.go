package main

import "log"

func main() {

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@smith.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Brown", "Anderson", "brown@smith.com", 17})
	users = append(users, User{"Alex", "Smith", "alex@smith.com", 44})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

	var firstLine = "once upon a midnight dreary"
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"
	animals["horse"] = "Charlie"

	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}

	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }
}

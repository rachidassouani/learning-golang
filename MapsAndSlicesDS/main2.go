package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myMap := make(map[string]User)

	me := User{
		FirstName: "Rachid",
		LastName:  "Assouani",
	}

	myMap["me"] = me

	log.Println("my first name is", myMap["me"].FirstName, "and my last name is", myMap["me"].LastName)

}

package main

import (
	"log"
	"time"
)

// struct like class in Java
type User struct {

	// In GoLang, variables that start with an uppercase letter mean that they are public, and if they start with a lowercase letter, they are private.
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	user := User{
		FirstName: "Rachid",
		LastName:  "Assouani",
	}

	log.Println(user.FirstName, user.LastName)
}

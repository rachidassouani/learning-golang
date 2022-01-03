package main

import "log"

func main() {

	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("Yes its true")
	}

	age := 30
	if age > 30 {
		log.Printf("Yes age is greater than 30")
	} else {
		log.Println("No, age is less than 30")
	}

	if age >= 30 && isTrue {
		log.Println("Yes, age is greater than 30, and isTrue var is true")
	}

}

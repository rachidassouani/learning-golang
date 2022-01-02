package main

import "log"

func main() {
	myMap := make(map[string]string)

	myMap["firstName"] = "Rachid"
	myMap["lastName"] = "Assouani"

	log.Println(myMap["firstName"])
	log.Println(myMap["lastName"])

}

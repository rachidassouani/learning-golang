package main

import "log"

func main() {
	for i := 0; i < 10; i++ {
		log.Println(i, "Hello")
	}

	// loop over a slice (Array)
	animals := []string{"dog", "cat", "horse", "cow", "fish"}
	for _, animal := range animals {
		log.Println(animal)
	}

	numbers := make(map[string]int)
	numbers["zero"] = 0
	numbers["one"] = 1
	numbers["two"] = 2

	// loop over a map
	for _, num := range numbers {
		log.Println(num)
	}

	// in case I want to show keys and values
	for key, num := range numbers {
		log.Println(key, num)
	}

}

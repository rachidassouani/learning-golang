package main

import (
	"log"
	"sort"
)

func main() {

	// slice of strings
	var names []string

	names = append(names, "Rachid")
	names = append(names, "Malika")

	log.Println(names)

	// slice of integers
	var numbers []int
	numbers = append(numbers, 1)
	numbers = append(numbers, 5)
	numbers = append(numbers, 8)
	numbers = append(numbers, 2)

	// we can sort them by using:
	sort.Ints(numbers)

	// printing the sorted slice
	log.Println(numbers)

	// other way to create slice
	ages := []int{10, 22, 35, 15}
	log.Println(ages)

	// what if want to print only the first 3 ages
	log.Println(ages[0:3])
}

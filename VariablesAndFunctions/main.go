package main

import "fmt"

func main() {
	// declaring a new STRING variable and printing it in the console.
	var name string
	name = "Rachid"
	fmt.Println(name)

	// declaring a new integer variable and printing it in the console.
	var age int
	age = 20
	fmt.Println(age)

	// we can also declare a variable and initialize it at the same time.
	var number = 5
	fmt.Println(number)

	// calling helloWorld function.
	printHelloWorld()

	// calling returnHello function and assigning the return data to a variable.
	var returnedData string = returnHello()
	fmt.Println(returnedData)

	// what if we don't know the return type.
	unknownDataType := returnHello()
	fmt.Println(unknownDataType)

	// calling function that return two strings.
	hello, world := returnHelloWorld()
	fmt.Println(hello, world)
}

// function that prints a string in the console.
func printHelloWorld() {
	fmt.Println("Hello, World.")
}

// function that return string.
func returnHello() string {
	return "Hello"
}

// function that return more than one item.
func returnHelloWorld() (string, string) {
	return "Hello", "world"
}

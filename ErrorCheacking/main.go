package main

import (
	"errors"
	"log"
)

func main() {
	firstInput := 55.5
	secondInput := 5

	result, error := devide(float32(firstInput), float32(secondInput))

	if error != nil {
		log.Println("Cannot devide by zero.")
		return
	}
	log.Println(result)
}

func devide(firstInput, secondInput float32) (float32, error) {
	if secondInput <= 0 {
		error := errors.New("Cannot devide by zero.")
		return 0, error
	}
	result := firstInput / secondInput
	return result, nil

}

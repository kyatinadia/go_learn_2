package main

import "fmt"

func main() {
	//initialization (memory address)
	// var firstNumber int = 4
	// var secondNumber *int = &firstNumber
	// fmt.Println()
	// fmt.Println("firstNumber (value): ", firstNumber)
	// fmt.Println("firstNumber (memory address): ", &firstNumber)
	// fmt.Println("secondNumber (value): ", *secondNumber)
	// fmt.Println("secondNumber (memory address): ", secondNumber)

	// var firstGroup string = "yeagerist"
	// var secondGroup *string = &firstGroup
	// fmt.Println("first group value: ", firstGroup)
	// fmt.Println("first group memory: ", &firstGroup)
	// fmt.Println("second group value: ", *secondGroup)
	// fmt.Println("second group memory: ", &secondGroup)

	// *secondGroup = "survey corps"
	// fmt.Println("first group value: ", firstGroup)
	// fmt.Println("first group memory: ", &firstGroup)
	// fmt.Println("second group value: ", *secondGroup)
	// fmt.Println("second group memory: ", &secondGroup)

	//pointer as parameter

	var age int = 10
	fmt.Println("Before:", age)
	changeValue(&age)
	fmt.Println("After:", age)

}

func changeValue(number *int) {
	*number = 20
}

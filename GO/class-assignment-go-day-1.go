// Section 1A: Variable Declaration and Data Types

package main

import (
	"fmt"
	"strings"
)

func class1() {
	// Variable Declaration:
	// Declare a variable named age and assign the value 25 to it.
	age := 25
	fmt.Println("age is :", age)

	// Data Types:
	// Define a variable name and assign your name as a string.
	name := "Sanket Kumar"
	fmt.Println("My Name is :", name)

	// Constants:
	// Create a constant named pi and assign the value 3.14 to it.
	const pi = 3.14
	fmt.Println("The PI vlaue is :", pi)

	// Type Conversion:
	// Convert the variable age to a float64 and store it in a new variable ageFloat

	ageFloat := float64(age)
	fmt.Printf("At first type of age was: %T and now the type of ageFloat is: %T\n", age, ageFloat)

	// Define a variable isSunny and assign a boolean value to it.
	isSunny := true

	fmt.Println(isSunny)

	// 	Section 1B: Additional Advanced Questions

	// Multiple Assignments:
	// Use multiple assignments to swap the values of two variables a and b without using a temporary variable.

	numberOne := 10
	numberTwo := 20

	fmt.Println(numberOne, numberTwo)

	numberOne = numberOne + numberTwo
	numberTwo = numberOne - numberTwo
	numberOne = numberOne - numberTwo

	fmt.Println(numberOne, numberTwo)

	// Anonymous Variables:
	// Ignore one of the return values of a function using an anonymous variable.

	result, _ := getData()
	fmt.Println("Result:", result)

	// Uninitialized Variables:
	// Declare a variable without initializing it and predict the default value assigned by Go.
	marks := 19.2 //float64
	fmt.Println(marks)

	// Custom Types:
	// Create a custom type Celcius for representing temperature in degrees Celsius. Implement methods to convert it to Fahrenheit and Kelvin.
	fmt.Println("The Fahrenheit is : ", getFahrenheit(37))
	fmt.Println("The Kelvin is :", getKelvin(37))

	// String Manipulation:
	// Declare a string variable with the value "Hello, Go!" and use string manipulation functions to print it in reverse.

	fmt.Println("Reversed string is :", reverseString("Hello, Go!"))

}

func getData() (int, string) {
	return 42, "Answer"
}

func getFahrenheit(degrees float64) float64 {
	return degrees*(1.8) + 32
}
func getKelvin(degrees float64) float64 {
	return degrees + 273
}

func reverseString(inputString string) string {
	fmt.Println("input String is :", inputString)
	var reverseStringed strings.Builder
	for i := len(inputString) - 1; i >= 0; i-- {
		reverseStringed.WriteByte(inputString[i])
	}
	return reverseStringed.String()
}

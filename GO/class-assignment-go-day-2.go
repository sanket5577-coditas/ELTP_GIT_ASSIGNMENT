// Section 2: Control Flow
package main

import (
	"fmt"
	// "strings"
)

func Day2() {
	// if-else Statements:
	// Write an if-else statement that prints "Even" if a variable num is even, and "Odd" if it's odd.

	inputNumber := 10

	if inputNumber%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// for Loop:
	// Use a for loop to print numbers from 1 to 5.

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Arrays:
	// Declare an array with integers 1, 2, 3, 4, and 5. Print each element using a for loop.
	fmt.Println("Array Question-----")
	numberArray := [5]int{1, 2, 3, 4, 5}

	for _, element := range numberArray {
		fmt.Println(element)
	}

	// Slices:
	// Create a slice from the array in the previous question and append the value 6 to it.
	mySlice := numberArray[:]
	mySlice = append(mySlice, 6)

	fmt.Println("MySlice is :", mySlice)

	// Maps:
	// Create a map representing a person with keys "name" and "age". Print the person's name.

	personMap := map[string]string{
		"name": "Sanket",
		"age":  "21",
	}

	fmt.Println("The PersonMap is :", personMap)

	// Nested Loops:
	// Use nested loops to print a pattern of asterisks resembling a pyramid.
	n := 5
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Switch with Expressions:
	// Write a switch statement that evaluates an expression and prints different messages based on the result.
	fmt.Println("Enter the day :")
	// var day string
	// fmt.Scan(&day)
	// day =strings.ToLower(day)
	// Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"
	// switch(day){
	// case "monday", "tuesday", "wednesday", "thursday", "friday":
	// 	fmt.Println("WeekDay")
	// case "saturday","sunday":
	// 	fmt.Println("WeekEnd")
	// default:
	// 	fmt.Println("Wrong Input Sorry,Try Again")
	// }

	// Break and Continue:
	// Use break to exit a loop prematurely and continue to skip the current iteration under a certain condition.

	for i := 1; i <= 10; i++ {
		if i == 8 {
			fmt.Println("Skipping 8")
			continue
		}

		fmt.Println(i)

		if i == 5 {
			fmt.Println("Breaking at 5")
			break
		}
	}
}

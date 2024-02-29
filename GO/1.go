package main

import "fmt"

func main() {
	//Variables
	// Declare
	var myName string
	var myAge int
	var isHappy bool

	// Printing the vaiables
	fmt.Println(myName)
	fmt.Println(myAge)
	fmt.Println(isHappy)

	// Declare + assigning values with datatypes
	var lastName string = "Kumar"
	fmt.Println(lastName)

	// Declare + assigning values 
	var employeeId = 5577
	fmt.Println(employeeId)

	//walrus Operator => short-hand 
	evenNumber := 2
	fmt.Println(evenNumber)

	floatNumber := 12.00;  //float64
	fmt.Println(floatNumber)

	// float=>float32,foat64
	// int => int8,int16,int32,int64

	characterLetters := 'a';
	fmt.Println(characterLetters);  //ascii value of a which is 97

	// samthing we can do with int32,the int32 is alias of rune
	var letterA int32 = 'a';
	fmt.Println(letterA)  //97


	// if we want the charater letter not the ascii value we can format the charaters using %

	fmt.Printf("%c",letterA)  //a
	fmt.Printf("%v",letterA)  //97
	fmt.Printf("%s",letterA)  //%!s(int32=97)


}

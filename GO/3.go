package main

import (
	"fmt"
	"strings"
)

func two() {
	message := "Hello"
	name := "Sanket"

	message = message + " " + name

	fmt.Println(message)
	length := len(message)
	fmt.Println("Length of Message is :", length)

	//String Methods

	indexOfIs := strings.Index(message, "ll")
	fmt.Println("index of ll is :", indexOfIs)
	indexOfIs = strings.Index(message, "xl")
	fmt.Println("index of ll is :", indexOfIs)

	contains := strings.Contains(message, "ll")
	fmt.Println(contains)

	contains = strings.Contains(message, "xl")
	fmt.Println(contains)

	userInput := "  sanket   "

	trimUserInput := strings.Trim(userInput, " ")
	fmt.Println(trimUserInput)

	trimUserInput = strings.Trim(trimUserInput, "s")
	fmt.Println(trimUserInput)

	//split

	csvInput := "sanket-21-5577"

	svInput := strings.Split(csvInput, "-")
	fmt.Println(svInput)

	stringName := "a|b|c|d"

	vara := strings.Replace(stringName, "b", "c", 1)

	fmt.Println(vara)
	vara = strings.ReplaceAll(stringName, "c", "b")

	fmt.Println(vara)


	word:="sanket"

	value:=word[0]

	fmt.Println(value)

	valueOfS := string(word[4])
	fmt.Println(valueOfS)

}

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func greet() {
	fmt.Println("Hello")
}
func Nine() {
	greet()

	ref := greet
	reference := ref

	reference()

	// fmt.Println(ref)

	add := func() {
		fmt.Println("i'm addd")

	}

	add()
	fmt.Println("addition is:", addition(2, 3))

	gotTheName:=GiveMeName("Sanket")

	fmt.Println("Hello",gotTheName)


	fmt.Println(RandomNumber())

}

func addition(a, b int) int {
	return a + b
}

func GiveMeName (name string) string{
	return strings.ToUpper(name)
}


func RandomNumber()(number,number2 float64){
	number = rand.Float64()
	number2 = rand.Float64()
	return 
}
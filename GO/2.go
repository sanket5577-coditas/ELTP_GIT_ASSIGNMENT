package main

import (
	"fmt"
	"math/rand"
)

func main(){
	const num1 int =2
	const num2 int =3

	const add = num1+num2
	const sub = num1-num2
	const div = num1/num2;
	const product = num1*num2

	// const div1 = 
	randomNumber := rand.Int31n(10)+30 
	fmt.Println(randomNumber)
}
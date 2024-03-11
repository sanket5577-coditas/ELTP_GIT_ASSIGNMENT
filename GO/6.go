package main

import (
	"fmt"
	
)

func Six() {
	var number [3]int
	numberss := make([]int, 3)
	numbersss := [3]int{}

	fmt.Println(number, numberss, numbersss)

	number[0] = 10
	number[1] = 20
	number[2] = 30

	numberss[0] = 1
	numberss[1] = 1
	numberss[2] = 1

	numbersss[0] = 2
	numbersss[1] = 2
	numbersss[2] = 2


	numberArray := [...]int{22,33,44,55}

	length := len(numberArray)

	fmt.Println(length)

	fmt.Println(number, numberss, numbersss)


	oddNumberArray := []int{1,3,5}
	evenNumberArray := []int{2,4,6}

	sumOfArray := append(oddNumberArray,evenNumberArray...)
	
	sumOfArray = append(sumOfArray,23131)
	fmt.Println(sumOfArray)

	   
}

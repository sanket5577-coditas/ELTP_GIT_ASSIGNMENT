package main

import "fmt"

func closure() {
	// count:=0
	// innerFunction := func()  int {
	// 	count++
	// 	return count
	// }
	// fmt.Println("Count:",innerFunction())
	// fmt.Println("Count:",innerFunction())
	// fmt.Println("Count:",innerFunction())

	functionTwo:=func (number int) int  {
		value:=number

		functionThree:=func () int {
			value++
			return value
		}
		return functionThree()
	}

	fmt.Println(functionTwo(5))
	fmt.Println(functionTwo(5))
	fmt.Println(functionTwo(5))
}
package main

import (
	"fmt"
	// "runtime/trace"
)

func name(){
	p:=1
	fmt.Println(p)
}
func name1(){
	p:=1
	fmt.Println(p)
}

func Four() {

	val := true && false
	fmt.Println(val)
	val = false && true
	fmt.Println(val)
	val = false && false
	fmt.Println(val)
	val = true && true
	fmt.Println(val) 

	fmt.Println()

	val = true || false
	fmt.Println(val)
	val = false || true
	fmt.Println(val)
	val = false || false
	fmt.Println(val)
	val = true || true
	fmt.Println(val)

	val = !true
	fmt.Println(val)

	val = !false
	fmt.Println(val)

	fmt.Println()
	num1 := 1
	num2 := 1.0

	isEqual := num1 == int(num2)
	fmt.Println(isEqual)
	fmt.Println("Other OPertor")
	isEqual = num1 >= int(num2)
	fmt.Println(isEqual)


	// mySlice:=[]int{1,2}
	// mySlice1:=[]int{1,2}

	// isEqual = mySlice>mySlice1
	// fmt.Println(isEqual)


	// myArray:=[2]int{1,2}
	// myArray1:=[2]int{1,2}

	// isEqual = myArray>myArray1

	// isEqual = name>name1

	str:="ancs"
	str1:="ancs"

	isEqual = str>str1
	fmt.Println(isEqual)

	// nam :=true
	// nam1 :=true
	// fmt.Println(nam>nam1)

}

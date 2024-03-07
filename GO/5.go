package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "msd is cool"
	age := 18

	if len(str) > 6 && strings.Contains(str, "cool") {
		fmt.Println("Sanket")
	} else if strings.Contains(str, "is") {
		fmt.Println("Sanket aniruddha")
	} else {
		fmt.Println("Aniruddh")
	}

	if age < 18 {
		fmt.Println("your kid ")
	} else if age > 18 && age < 25 {
		fmt.Println("Allowed on 2nd floor")
	} else if age > 25 {
		fmt.Println("first Floor")
	} else {
		fmt.Println("Not allowed")
	}


	
}

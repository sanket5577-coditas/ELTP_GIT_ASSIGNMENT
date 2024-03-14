package main

import "fmt"

func Eight() {
	person := map[string]string{}

	fmt.Println(person)


	person["Name"] = "Sanket"
	person["Age"]  = "21"

	fmt.Println(person)

	delete(person,"Age")

	fmt.Println(person)


	bookDetail:=map[string]string{
		"author":"Saket",
		"publisher":"Coditas",
	}

	fmt.Println(bookDetail)

	bookDetail2:=map[string]string{
		"author":"Saket",
		"publisher":"Coditas",
	}

	fmt.Println(bookDetail2)

	// map with an array as the key
    myMap := make(map[[3]int]string)

    // Create an array to use as a key
    key := [3]int{1, 2, 3}

    // Assign a value to the map using the array key
    myMap[key] = "Value associated with the key array"

    // Retrieve the value from the map using the same array key
    value := myMap[key]

    // Print the retrieved value
    fmt.Println(value) // Output: Value associated with the key array
}

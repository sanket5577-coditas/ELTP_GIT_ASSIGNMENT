package main

import "fmt"

func main() {
	n := 5

	// 1
	// 1 2
	// 1 2 3
	// 1 2 3 4
	// 1 2 3 4 5

	// for i := 1; i <=n; i++ {
	// 	for j := 1; j <=i; j++ {
	// 		fmt.Print(j," " )
	// 	}
	// 	fmt.Println()
	// }

	//		 1
	// 		1 2
	// 	 1 2 3
	// 1 2 3 4
	//1 2 3 4 5

	// for i := 1; i <=n; i++ {
	// 	for j := 1; j <=n-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for j:=1; j<=i; j++ {
	// 		fmt.Print(j)
	// 	}
	// 	fmt.Println()
	// }

	sanket := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j1 := 1; j1 <= i; j1++ {
			fmt.Print(sanket, "")
			sanket++

		}
		sanket = sanket
		fmt.Println()
		
	}
}

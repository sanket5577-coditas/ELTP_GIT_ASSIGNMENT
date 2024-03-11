package main

import "fmt"

func Pattern() {
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

	// 	 1
	// 	232
	//    34543
	//   4567654
	//  567898765

	// 	for i := 1; i <= n; i++ {
	// 		for j := 1; j <= n-i; j++ {
	// 			fmt.Print(" ")
	// 		}
	// 		for j1 := i; j1 <= 2*i-1; j1++ {
	// 			fmt.Print(j1, "")

	// 		}
	// 		for j := 2*i - 2; j >= i; j-- {
	// 			fmt.Print(j)
	// 		}

	// 		fmt.Println()

	// 	}

	// 54321
	//  4321
	//   321
	//    21
	//     1

	// for i := n; i >= 1; i-- {

	// 	for j := n - i; j > 0; j-- {
	// 		fmt.Print(" ")
	// 	}

	// 	for j := i; j >= 1; j-- {
	// 		fmt.Print(j)
	// 	}

	// 	fmt.Println()
	// }

	// for i := 1; i <= n; i++ {
	// 	for j := i; j <= n; j++ {
	// 		if i==1 || i==n {
	// 			fmt.Print("*")
	// 		} else if j==1 || j==n{
	// 			fmt.Println("*")
	// 		}else {
	// 			fmt.Print(" ")
	// 		}
	// 	}

	// 	fmt.Println()
	// }
}

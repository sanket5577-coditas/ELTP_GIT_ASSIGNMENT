package main

import "fmt"

func getMonthlyPrice(tier string) int {
    if tier == "basic" {
        return 100
    } else if tier == "premium" {
        return 150
    } else if tier == "enterprise" {
        return 500
    } else {
        return 0
    }
}

func Func01(){
	fmt.Println(getMonthlyPrice("basic"))
}






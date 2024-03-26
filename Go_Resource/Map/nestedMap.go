package main

// import (
// 	"fmt"
// )

// func getNameCounts(names []string) map[rune]map[string]int {
// 	nameCounts := make(map[rune]map[string]int)
// 	for _, name := range names {
// 		firstLetter := rune(0)
// 		if len(name) != 0 {
// 			firstLetter = rune(name[0])
// 		}

// 		if _, ok := nameCounts[firstLetter]; !ok {
// 			nameCounts[firstLetter] = make(map[string]int)
// 		}
// 		if _, ok := nameCounts[firstLetter][name]; !ok {
// 			nameCounts[firstLetter][name] = 0
// 		}
// 		nameCounts[firstLetter][name]++
// 	}
// 	return nameCounts
// }


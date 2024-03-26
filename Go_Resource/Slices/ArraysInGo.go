package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {

	messages := [3]string{primary, secondary, tertiary}
	costs := [3]int{}

	costs[0] = len(primary)

	if secondary != "" {
		costs[1] = costs[0] + len(secondary)
	}

	if tertiary != "" {
		costs[2] = costs[1] + len(tertiary)
	}

	return messages, costs
}

func main() {
	primary := "hello"
	secondary := "world"
	tertiary := "!"

	messages, costs := getMessageWithRetries(primary, secondary, tertiary)

	fmt.Println( messages)
	fmt.Println(costs)
}

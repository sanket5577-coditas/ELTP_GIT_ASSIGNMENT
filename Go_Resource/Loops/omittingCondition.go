package main

func maxMessages(thresh int) int {
	totalCost := 0.0

	for i := 0;;i++ {
		totalCost += 1.01 + (0.01 * float64(i))
	}
	// if totalCost > thresh {
	// 	return 1
	// }
}

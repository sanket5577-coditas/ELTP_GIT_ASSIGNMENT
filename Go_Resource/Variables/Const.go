package main

import "fmt"

func Const() {
	const premiumPlanName = "Premium Plan"
	// premiumPlanName = "Basic Plan"
	const basicPlanName ="Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)

	///Other Question code

	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute*minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}



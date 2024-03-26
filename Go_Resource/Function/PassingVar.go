package main
import "fmt"

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}

func Pass(){

	costPerSend := 200 
    numLastMonth := 800 
    numThisMonth := 940 
	increase := monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth)
   
    fmt.Println("Monthly bill increase:", increase)
}
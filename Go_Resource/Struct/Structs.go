package main

import "fmt"	



func main(){
	 type messageToSend struct {
		phoneNumber int
		message string
	}

	messageTosend:=messageToSend{4949,"ncn"}
	
fmt.Println(messageTosend.message)

// Gauri:= map[string]string{
// 	"Gender":"Male",
// }

// name,found:=Gauri["Name"]
// fmt.Println(name,found)

}
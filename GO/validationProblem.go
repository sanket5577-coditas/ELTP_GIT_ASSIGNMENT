package main



// import "fmt"

// func isLengthValid(number int) func(input string) bool {
// 	compare := func(input string) bool {
// 		return len(input) > number
// 	}
// 	return compare
// }

// func isString(input string) bool {

// 	for _, char := range input {
// 		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
// 			return false
// 		}
// 	}
// 	return true
// }

// func validation(input string, fns... func(string) bool) bool {
// 	for _,validate:= range fns {
// 		if fns[validate] == true{
// 			return true
// 		}
// 	}
// 	return false
// }


func isNumber(input int) bool{
	if input>5 {
		return true
	}
	return false
}

func isEven(input int) func (input int) bool{
	istrue:= func(input int) bool{
		if input % 2 == 0{
			return true
		}
		return false
	}
	return istrue

}

func validationPipeline(input int , fns... func(input int) bool)bool {
	for numbers := range len(fns){
		fn:=fns[numbers]
		if !fn(input) {
			return true
		}
	}
	return false
}
func main() {
	input := 5
	// validationPipeline := []func(input int) bool{isNumber(input),isEven}

	validationPipeline(input,isNumber,isEven(input))

}

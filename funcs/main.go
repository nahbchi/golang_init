package main

import ("fmt"
 "errors"
)
func main() {
	var PrintValue string = "Hello"
	PrintMe(PrintValue)
	

	var numerator int =  11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
    if err!=nil{
		fmt.Printf(err.Error())
	}else if remainder == 0{
		fmt.Printf("The result of the integer division is %v", result)
	}else{
		fmt.Printf("The result of the integer division is %v with remaninder %v", result, remainder)
	}
	
}

func PrintMe(PrintValue string) {
	fmt.Println("Hello World")
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator ==0{
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
 var result int = numerator/denominator
 var remainder int = numerator%denominator
 return result, remainder, err
}
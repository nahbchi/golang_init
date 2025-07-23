package main

import "fmt"
import "unicode/utf8"
func main(){
	var intNum int 
	fmt.Println(intNum)

	var floatNum float64
	fmt.Println(floatNum)

	var intNum1 int = 3
	var intNum2 int =  2
	fmt.Println(intNum2/intNum1)
	fmt.Println(intNum1%intNum2)

	var MyString string = "Hello"+ " " + "World"
	fmt.Println(MyString)
	
	fmt.Println(len("Test"))
    fmt.Println(utf8.RuneCountInString("hallo"))

	var Myboolean bool = false 
	fmt.Println(Myboolean)

	var IntNum3 rune
	fmt.Println(IntNum3)

	 myVar := "text"	
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const MyConst string = "const value"
	fmt.Println(MyConst)

	const pi float32 = 3.1415 
}
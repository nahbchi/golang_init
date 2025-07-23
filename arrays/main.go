package main

import (
	"fmt"

)

func main(){
	var intArr =[...]int32{1,2,3}
	var i int = 0
	fmt.Println(intArr)
	
	var IntSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v", len(IntSlice), cap(IntSlice))
	IntSlice = append(IntSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(IntSlice), cap(IntSlice))
  
    var IntSlice2 []int32 = []int32{8,9}
	IntSlice = append(IntSlice, IntSlice2...)
	fmt.Println(IntSlice)

	var MyMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(MyMap)

	var MyMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(MyMap2["Adam"])
	var age, ok = MyMap2["Jason"]
	
	if ok{
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Printf("Invalid Name")
	}

	for name := range MyMap2{
		fmt.Printf("\nName: %v, Age: %v\n", name, age)
	}
	
	for i, v := range intArr{
		fmt.Printf("\nIndex: %v, Value: %v\n", i, v)
	}

	for i=0; i<10; i++{
		fmt.Println(i)
	}
}

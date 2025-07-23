package main

import (
	"fmt"	
	"strings"
)

func main(){
	var Mystring =[]rune("resume")
	var Indexed = Mystring[0]
	fmt.Printf("%v,%T", Indexed, Indexed)
	for i, v := range Mystring{
		fmt.Println(i, v)
}
  fmt.Printf("the length of 'MyString' is %v", len(Mystring))

  var MyRune = 'a'
  fmt.Printf("\nMyrune = %v", MyRune)

  var StrSlice = []string{"n", "a", "h", "b", "c", "h", "i"}
  var StrBuilder strings.Builder
  for i := range StrSlice{
	StrBuilder.WriteString(StrSlice[i])
  }
  var CatStr = StrBuilder.String()
  fmt.Printf("\n%v", CatStr)

}
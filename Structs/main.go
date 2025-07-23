package main
import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
	owner
}

type eletricEngine struct{
	mpkwh uint8
	kwh uint8
}
type owner struct{
	name string
}


func(e eletricEngine) milesleft() uint8{
	return e.kwh*e.mpkwh
}

func(e gasEngine) milesleft() uint8{
	return e.gallons*e.mpg
}

type engine interface{
	milesleft() uint8
}

func CanMakeIt(e engine, miles uint8){
	if miles<=e.milesleft(){
		fmt.Println("You Can make it there!")
	}else{
		fmt.Println("Need to fuel the tank first!")
	}
}
func main(){
  var MyEngine gasEngine = gasEngine{25, 15, owner{"Jhon"}}
  CanMakeIt(MyEngine,50)
 }
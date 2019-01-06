package main
import (
	"fmt"
)


func main(){
	// var num1 int32 = 200
	// var num2 int32 = 30
	// var num3 int32 = num1 / num2
	// var b bool = num1 > num2
	// fmt.Println(num3)
	// fmt.Println(b)
	var i int = 10
	// if i < 9 ＆＆ true {
	// 	fmt.Println("ok")
	// }
	i += 3
	fmt.Println(i)
	num1 := 2
	num2 := 3
	// num := num1
	// num1 = num2
	// num2 = num
	
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println(num1, num2)

	a := 2
	b := 1
	c := 1
	numMax := a
	if b > numMax {
		numMax = b
	}
	if c > numMax {
		numMax = c
	}
	fmt.Println(numMax)
}
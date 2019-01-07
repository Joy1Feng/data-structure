package main
import (
	"fmt"
)


func myCal (num1 int32, num2 int32) (sum int32, sub int32) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main () {
	var num1 int32 = 10
	var num2 int32 = 10
	num3, num4 := myCal(num1, num2)
	fmt.Println(num3, num4)
}

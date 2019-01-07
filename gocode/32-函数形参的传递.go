package main
import (
	"fmt"
)


// func cal (num1 *int32) {
// 	*num1++
// 	fmt.Println(num1)
// }

// func cal1 (funcvar func (*int32), num1) {
// 	funcvar(num1)
// }

type myFun func (int32, int32) int32

func cal (num1 int32, num2 int32) int32 {
	return num1 + num2
}

func cal1 (funcvar myFun, num1 int32, num2 int32) int32 {
	return funcvar(num1, num2)
}


func main () {
// 	var num1 int32 = 10
// // 	cal(&num1)
// // 	fmt.Println(num1)
// 	cal1(cal, num1)
// 	fmt.Println(num1)
	res := cal1(cal, 10 ,10)
	fmt.Println(res)
}
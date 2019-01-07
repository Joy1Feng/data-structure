package main
import (
	"fmt"
)

func fib (num int32) int32 {
	if num == 1 || num == 2 {
		return 1
	} else {
		return fib(num -1) + fib(num - 2)
	}
	
}

func cal (num int32) int32 {
	if num == 1 {
		return 3
	} else {
		return 2 * cal(num - 1) + 1
	}
}

func main () {
	var num int32 = 10
	res := fib(num)
	fmt.Println(res)
	res1 := cal(num)
	fmt.Println(res1)
}
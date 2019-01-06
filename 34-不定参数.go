package main
import (
	"fmt"
)

func init () {
	fmt.Println("init被调用")
}

func cal (num1 int, args... int) (sum int) { //sum int这个括号必须有
	sum = num1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func cal1 (args... int) int {
	var sum int = 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main () {
	var num1 int = 1
	var num2 int = 2
	var num3 int = 3
	res := cal(num1, num2, num3)
	fmt.Println(res)
	res1 := cal1(1, 2, 3)
	fmt.Println(res1)
}
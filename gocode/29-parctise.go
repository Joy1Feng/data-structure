package main
import (
	"fmt"
)

func cal(num1 int64, num2 int64) (int64, int64) {
	var res1 int64
	var res2 int64
	res1 = num1 + num2
	res2 = num1 - num2
	return res1, res2
}

func main(){
	res1, _ := cal(1, 2)
	fmt.Println(res1)
}
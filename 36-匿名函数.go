package main
import (
	"fmt"
)


func main () {
	//法1 匿名函数自调
	res := func (num1 int, num2 int) int {
		return num1 + num2
	} (1, 2)
	fmt.Println(res)

	//法2　匿名函数自调
	a := func (num1 int, num2 int) int {
		return num1 + num2
	}
	res1 := a (1, 2)
	res2 := a (1, 2)
	fmt.Println(res1, res2)
}



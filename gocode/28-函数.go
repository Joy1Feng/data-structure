package main
import (
	"fmt"
)

func cal(n1 float32, n2 float32, oprator byte) float32 {
	var res float32
	switch oprator {
		case '+':
			res = n1 + n2
			// fmt.Printf("%f + %f = %f\n", n1, n2, res)
		case '-':
			res = n1 - n2
			// fmt.Printf("%f － %f = %f\n", n1, n2, res)
		case '*':
			res = n1 * n2
			// fmt.Printf("%f ＊ %f = %f\n", n1, n2, res)
		case '/':
			res = n1 / n2
			// fmt.Printf("%f ／ %f = %f\n", n1, n2, res)
		default:
			fmt.Println("你输入的内容有误!")
	}
	return res
}

func main(){
	var op byte = '+'
	res1 := cal(1.2, 1.4, op)
	fmt.Println(res1)
}
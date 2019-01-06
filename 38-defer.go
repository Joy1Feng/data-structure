package main
import (
	"fmt"
)

func sum (num1 int, num2 int) int {
	//当程序执行到defer时,暂时不执行,会将defer后面的语句压入到独立的defer栈,
	//当函数执行完毕后,再从defer栈,按照先入后出的出栈,执行
	//在入栈的时候也会把相关的值拷贝入栈
	defer fmt.Println("n1=", num1)
	defer fmt.Println("n2=", num2)
	fmt.Println("sum=", num1 + num2)
	num1++
	num1++
	return num1 + num2
}

func main () {
	res := sum(1, 2)
	fmt.Println(res)
}